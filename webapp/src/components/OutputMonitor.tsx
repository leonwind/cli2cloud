import {Component} from "react";
import { useParams } from "react-router-dom";
import * as grpcWeb from "grpc-web";
import {Cli2CloudClient} from "../proto/ServiceServiceClientPb"
import {Client, Content} from "../proto/service_pb"



const cli2CloudService = new Cli2CloudClient("http://localhost:8000", null, null)

export function OutputMonitor() {
    let params = useParams();
    const client = new Client();

    if (params.clientID === undefined) {
        return <div>Client ID is undefined</div>;
    }

    client.setId(params.clientID)
    
    const stream = cli2CloudService.subscribe(client, {});
    stream.on("data", function(response: Content) {
        console.log("Hallo");
        console.log(response.getPayload(), response.getRow());
    });
    stream.on("error", (error: Error): void => {
        console.error(error);
    });

    console.log("HERE I AM");
        
    return <div>{params.clientID}</div>;
}

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

        this.state = {
            clientID: "3Q8N0o",
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
            //console.log(response.getPayload(), response.getRow());
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
        const allRows: JSX.Element[] = this.state.contents.map((row: Row) => 
            <div>Row {row.line}, content: {row.content}</div>
        );
        console.log(this.state.contents.length);
                
        return (
            <div> 
                {allRows}
            </div>
        );
    }
}