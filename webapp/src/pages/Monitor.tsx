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
    decryptor: DecryptionService | null,
    rows: Row[],
    raw: boolean,
}

export class Monitor extends Component<{}, State> {
    private numLines: number;
    private cli2CloudService: Cli2CloudClient;
    private clientId: ClientId;
    private client: Promise<Client>;

    constructor(props: any) {
        super(props);

        this.state = {
            encrypted: false,
            decryptor: null,
            enterPwdFirstTime: true,
            rows: [],
            raw: new URLSearchParams(new URL(window.location.href).search).has("raw"),
        };

        this.numLines = 1;
        this.cli2CloudService = new Cli2CloudClient("https://cli2cloud.com:1443", null, null);

        const id = window.location.pathname.substring(1);
        
        this.clientId = new ClientId();
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
        this.loadContent();
    }

    private updatePassword(newPassword: string) {
        this.createDecryptor(newPassword);
    }

    private createDecryptor(password: string) {
        this.client.then((client: Client) => {
            this.setState({decryptor: new DecryptionService(password, client.getSalt(), client.getIv())});
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
        window.location.hash = line.toString();
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
        return this.state.rows.map((row: Row) => 
            <div className={styles.row} id={row.line.toString()} key={row.line}>
                <span className={styles.line} onClick={() => this.highlightRow(row.line)}>
                    {row.line}
                </span>
                <span className={styles.content}>
                    {this.decryptRowIfEncrypted(row.content)}
                </span>
            </div>
        );
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

    private switchToRawData() {
        let params = new URLSearchParams(new URL(window.location.href).search);
        params.set("raw", "true");
        window.location.search = params.toString()

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