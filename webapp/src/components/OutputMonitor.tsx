import { useParams } from "react-router-dom";
import {grpc} from "@improbable-eng/grpc-web";



export function OutputMonitor() {
    let params = useParams();
            
    return <div>{params.clientID}</div>;
}
