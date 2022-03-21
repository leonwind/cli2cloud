import {Component} from "react";
import styles from "../styles/Monitor.module.css";
import {Cli2CloudClient} from "../proto/ServiceServiceClientPb"
import {Client, ClientId, Payload} from "../proto/service_pb"
import {DecryptionService} from "../services/DecryptionService"
import {NavBar} from "../components/NavBar"
import { ChangeDecryptionPwd } from "../components/ChangeDecryptionPwd";


interface Row {
    content: string,
    line: number,
}

interface State {
    encrypted: boolean,
    enterPwdFirstTime: boolean,
    password: string | null,
    decryptor: DecryptionService | null,
    rows: Row[],
    raw: boolean,
    highlightRow: string,
}

export class Monitor extends Component<{}, State> {
    private numLines: number;
    private cli2CloudService: Cli2CloudClient;
    private clientId: ClientId;
    private client: Promise<Client>;

    constructor(props: any) {
        super(props);

        // Redirect due to backward compatibility with old client which set the
        // key as a query parameter(?key=) and not as a hash parameter (#key=)
        let params = new URLSearchParams(new URL(window.location.href).search);
        if (params.has("key")) {
            const password = params.get("key");
            if (password !== null) {
                this.addToHashParam("key", password);
            }
            params.delete("key");
            window.location.search = params.toString()
        }
        
        let password = this.extractFromHash(window.location.hash, "key");
        let highlightRowId = this.extractFromHash(window.location.hash, "row");

        this.state = {
            encrypted: false,
            enterPwdFirstTime: password === null,
            password: password,
            decryptor: null,
            rows: [],
            raw: params.has("raw"),
            highlightRow: highlightRowId === null ? "" : highlightRowId,
        };

        this.numLines = 1;
        //this.cli2CloudService = new Cli2CloudClient("https://cli2cloud.com:1443", null, null); // production
        this.cli2CloudService = new Cli2CloudClient("http://localhost:8000", null, null); // local dev

        this.clientId = new ClientId();
        const id = window.location.pathname.substring(1);
        this.clientId.setId(id);

        this.client = this.cli2CloudService.getClientById(this.clientId, {})
    }

    componentDidMount() {
        this.loadContent = this.loadContent.bind(this);
        this.highlightRow = this.highlightRow.bind(this);
        this.updatePassword = this.updatePassword.bind(this);
        this.afterFirstTimePassword = this.afterFirstTimePassword.bind(this);
        this.switchToRawData = this.switchToRawData.bind(this);

        this.client.then((client) => {this.setState({encrypted: client.getEncrypted()})});

        if (!this.state.enterPwdFirstTime) {
            this.createDecryptor(this.state.password);
        }

        this.loadContent();
    }

    private extractFromHash(hash: string, key: string): string | null {
        const params: string = hash.substring(1, hash.length);
        let value: string | null = null;

        params.split("&").forEach((parts, _) => {
            let kv = parts.split("=");
            if (kv !== [] && kv[0] === key) {
                value = kv[1];
            }
        });
        return value;
    }

    private addToHashParam(key: string, value: string, remove: boolean=false) {
        const newParamPair = key + "=" + value;
        const currHash = window.location.hash.substring(1, window.location.hash.length);
        let newHash = "";
        let exists: boolean = false;

        currHash.split("&").forEach((parts, _) => {
            let kv = parts.split("=");
            if (kv.length !== 0 && kv[0] !== '') {
                if (kv[0] === key) {
                    exists = true;
                    if (remove) {
                        return;
                    }
                    newHash += newParamPair;
                } else {
                    newHash += parts;
                }
                newHash += '&';
            }
        });
        
        if (!exists) {
            newHash += newParamPair;
        }

        window.location.hash = newHash;
    }

    private updatePassword(newPassword: string) {
        this.addToHashParam("key", newPassword);
        this.setState({password: newPassword});
        this.createDecryptor(newPassword);
    }

    private createDecryptor(password: string | null) {
        if (password === null) {
            console.log("Can't create decryptor");
            return;
        }
        this.client.then((client: Client) => {
            this.setState({decryptor: new DecryptionService(password!, client.getSalt(), client.getIv())});
        });
    }

    private loadContent() {
        const stream = this.cli2CloudService.subscribe(this.clientId, {});

        stream.on("data", (response: Payload) => {
            this.addNewContent(response.getBody())
        });

        stream.on("error", (error: Error): void => {
            console.error(error);
        });

        
    }

    private addNewContent(content: string) {
        let newRows: Row[] = this.state.rows;
        newRows.push({
            content: content,
            line: this.numLines,
        });

        this.numLines += 1
        this.setState({rows: newRows});
    } 

    private highlightRow(line: number) {
        if (this.state.highlightRow === line.toString()) {
            this.setState({highlightRow: ""});
            // delete the hash parameter again if set
            this.addToHashParam("row", "", true);
        } else {
            this.addToHashParam("row", line.toString());
            this.setState({highlightRow: line.toString()});
        }
    }

    private decryptRowIfEncrypted(content: string): string {
        if (this.state.encrypted && this.state.decryptor !== null) {
            return this.state.decryptor.decrypt(content);
        }
        return content; 
    }

    private createNewDecryptorIfEncrypted() {
        // Since we decrypt everything again from the beginning, 
        // we need to init the decryptor from the beginning as well.
        if (this.state.decryptor !== null) {
            this.state.decryptor.createDecryptor();
        }
    }

    private createDivsForAllRows(): JSX.Element[] | JSX.Element {
        if (this.state.rows.length === 0) {
            return [<div className={styles.emptyRows}>
                No output found for client "{this.clientId.getId()}".
            </div>];
        }

        this.createNewDecryptorIfEncrypted() 
        return this.state.rows.map((row: Row) => {
            let rowStyle = row.line.toString() === this.state.highlightRow ? styles.selectedRow : styles.row;

            return <div className={rowStyle} id={row.line.toString()} key={row.line}>
                <span className={styles.line} onClick={() => this.highlightRow(row.line)}>
                    {row.line}
                </span>
                <span className={styles.content}>
                    {this.decryptRowIfEncrypted(row.content)}
                </span>
            </div>
        });
    }

    private createDivsForRawOutput(): JSX.Element[] | JSX.Element {
        if (this.state.rows.length === 0) {
            return <div>No output found for client "{this.clientId.getId()}."</div>
        }

        this.createNewDecryptorIfEncrypted()
        return this.state.rows.map((row: Row) => 
            <div key={row.line}>{this.decryptRowIfEncrypted(row.content)}</div> 
        );
    }

    private setURLParams(key: string, value: string) {
        let params = new URLSearchParams(new URL(window.location.href).search);
        params.set(key, value);
        window.location.search = params.toString()
    }

    private switchToRawData() {
        this.setURLParams("raw", "true");
        this.setState({raw: true});
    }

    private afterFirstTimePassword() {
        this.setState({enterPwdFirstTime: false});
    }

    render() {
        if (this.state.raw) {
            return this.createDivsForRawOutput()
        }

        return (
            <>
                {this.state.encrypted && this.state.decryptor === null &&
                <ChangeDecryptionPwd show={this.state.enterPwdFirstTime} onSubmit={this.updatePassword} onClose={this.afterFirstTimePassword}/>}

                <NavBar showPasswordBtn={this.state.encrypted} onPasswordSubmit={this.updatePassword} switchToRawData={this.switchToRawData}/>
                <div className={styles.body}>
                    <div className={styles.outputArea}>
                        {this.createDivsForAllRows()}
                    </div>
                </div>
            </>
        );
    }
}