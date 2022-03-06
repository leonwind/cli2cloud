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
    decryptor: DecryptionService | null,
    inputPassword: boolean,
    contents: Row[],
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
            inputPassword: true,
            contents: [],
        };
        
        this.numLines = 0;
        this.cli2CloudService = new Cli2CloudClient("http://localhost:8000", null, null)

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
            this.addNewContent(this.decryptRowIfEncrypted(response.getBody()));
        });

        stream.on("error", (error: Error): void => {
            console.error(error);
        });
    }

    private addNewContent(content: string) {
        let new_content: Row[] = this.state.contents;
        new_content.push({
            content: content,
            line: this.numLines,
        });

        this.numLines += 1
        this.setState({contents: new_content});
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

    private createDivsForAllRows(): JSX.Element[] {
        return this.state.contents.map((row: Row) => 
            <div className={styles.row} id={row.line.toString()} key={row.line}>
                <span className={styles.line} onClick={() => this.highlightRow(row.line)}>
                    {row.line}
                </span>
                <span className={styles.content}>
                    {row.content}
                </span>
            </div>
        );
    }

    private afterFirstTimePassword() {
        this.setState({inputPassword: false});
    }

    render() {
        let allRows: JSX.Element[];

        if (this.state.contents.length === 0) {
            allRows = [<div>No output for client ID "{this.clientId.getId()}".</div>];
        } else {
            allRows = this.createDivsForAllRows();
        }

        return (
            <>
                {this.state.encrypted && this.state.decryptor === null &&
                <ChangeDecryptionPwd show={this.state.inputPassword} onSubmit={this.updatePassword} onClose={this.afterFirstTimePassword}/>}

                <NavBar showPasswordBtn={this.state.encrypted} onPasswordSubmit={this.updatePassword}/>
                <div className={styles.body}>
                    <div className={styles.outputArea}>
                        {allRows}
                    </div>
                </div>
            </>
        );
    }
}