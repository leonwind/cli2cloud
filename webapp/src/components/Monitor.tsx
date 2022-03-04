import {Component} from "react";
import styles from "../styles/Monitor.module.css";
import {Cli2CloudClient} from "../proto/ServiceServiceClientPb"
import {Client, ClientId, Payload} from "../proto/service_pb"
import {DecryptionService} from "../services/DecryptionService"


interface Row {
    content: string,
    line: number,
}

interface State {
    contents: Row[],
    encrypted: boolean,
    decryptor: DecryptionService | null,
}

export class Monitor extends Component<{}, State> {
    private numLines: number;
    private cli2CloudService: Cli2CloudClient;
    private client: Promise<Client>;
    private clientId: ClientId;

    constructor(props: any) {
        super(props);

        this.state = {
            contents: [],
            encrypted: true,
            decryptor: new DecryptionService("123", "", ""),
        };
        
        this.numLines = 0;
        this.cli2CloudService = new Cli2CloudClient("http://localhost:8000", null, null)
        this.clientId = new ClientId();
        this.clientId.setId(window.location.pathname.substring(1));        
        this.client = this.cli2CloudService.getClientById(this.clientId, {});
        console.log(this.clientId);
        console.log(this.client);

        this.loadContent = this.loadContent.bind(this);
        this.addNewContent = this.addNewContent.bind(this);
        this.createDivsForAllRows = this.createDivsForAllRows.bind(this);
        this.highlightRow = this.highlightRow.bind(this);
    }

    componentDidMount() {
        this.loadContent();
    }

    private loadContent() {
        const stream = this.cli2CloudService.subscribe(this.clientId, {});

        stream.on("data", async (response: Payload) => {
            if ((await this.client).getEncrypted()) {
                this.addNewContent(response.getBody());    
            } else {
                this.addNewContent(response.getBody());
            }
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

    private createDivsForAllRows(): JSX.Element[] {
        return this.state.contents.map((row: Row) => 
            <li className={styles.row} id={row.line.toString()} key={row.line}>
                <span className={styles.line} onClick={() => this.highlightRow(row.line)}>{row.line}</span>
                <span className={styles.content}>{row.content}</span>
            </li>
        );
    }

    render() {
        let allRows: JSX.Element[];

        if (this.state.contents.length === 0) {
            allRows = [<div>No output for client ID "{this.clientId.getId()}".</div>];
        } else {
            allRows = this.createDivsForAllRows();
        }

        return (
            <div className={styles.body}>
                <div className={styles.outputArea}>
                    {allRows}
                </div>
            </div>
        );
    }
}