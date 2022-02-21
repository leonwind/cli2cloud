import {Component} from "react";
import { useParams } from "react-router-dom";
import * as grpcWeb from "grpc-web";
import {Cli2CloudClient} from "../proto/ServiceServiceClientPb"
import {Client, Content} from "../proto/service_pb"

const cli2CloudService = new Cli2CloudClient("http://localhost:8000", null, null)

interface Row {
    content: string,
    line: number,
}

interface State {
    clientID: string,
    contents: Row[],
}

export class Monitor extends Component<{}, State> {

    constructor(props: any) {
        super(props);

        const clientID = window.location.pathname.substring(1);

        this.state = {
            clientID: clientID,
            contents: []
        };
        this.loadContent = this.loadContent.bind(this);
        this.addNewContent = this.addNewContent.bind(this);        
    }

    componentDidMount() {
        this.loadContent();
    }

    private loadContent() {
        const client = new Client();
        client.setId(this.state.clientID)

        const stream = cli2CloudService.subscribe(client, {});

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

    render() {
        if (this.state.contents.length === 0) {
            return <div>No output for client ID "{this.state.clientID}".</div>;
        }

        const allRows: JSX.Element[] = this.state.contents.map((row: Row) => 
            <div>Row {row.line}, content: {row.content}</div>
        );

        return (
            <div>{allRows}</div>
        );
    }
}