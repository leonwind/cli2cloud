import {Component} from "react";
import styles from "../styles/Monitor.module.css";
import {Cli2CloudClient} from "../proto/ServiceServiceClientPb"
import {Client, Content} from "../proto/service_pb"


interface Row {
    content: string,
    line: number,
}

interface State {
    contents: Row[],
}

export class Monitor extends Component<{}, State> {
    private cli2CloudService: Cli2CloudClient;
    private clientID: string;

    constructor(props: any) {
        super(props);

        this.state = {
            contents: []
        };

        this.clientID = window.location.pathname.substring(1);
        this.cli2CloudService = new Cli2CloudClient("http://localhost:8000", null, null)

        this.loadContent = this.loadContent.bind(this);
        this.addNewContent = this.addNewContent.bind(this);
        this.createDivsForAllRows = this.createDivsForAllRows.bind(this);
        this.highlightRow = this.highlightRow.bind(this);
    }

    componentDidMount() {
        this.loadContent();
    }

    private loadContent() {
        const client = new Client();
        client.setId(this.clientID)

        const stream = this.cli2CloudService.subscribe(client, {});

        stream.on("data", (response: Content) => {
            this.addNewContent(response);
        });

        stream.on("error", (error: Error): void => {
            console.error(error);
        });
    }

    private addNewContent(new_row: Content) {
        let new_content: Row[] = this.state.contents;
        new_content.push({
            content: new_row.getPayload(), 
            line: new_row.getRow()
        });

        this.setState({contents: new_content});
    } 

    private highlightRow(line: number) {
        window.location.hash = line.toString();
    }

    private createDivsForAllRows(): JSX.Element[] {
        return this.state.contents.map((row: Row) => 
            <div className={styles.row} id={row.line.toString()} key={row.line}>
                <span className={styles.line} onClick={() => this.highlightRow(row.line)}>{row.line}</span>
                <span className={styles.content}>{row.content}</span>
            </div>
        );
    }

    render() {
        let allRows: JSX.Element[];

        if (this.state.contents.length === 0) {
            allRows = [<div>No output for client ID "{this.clientID}".</div>];
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