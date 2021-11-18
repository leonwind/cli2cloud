import { useParams } from "react-router-dom";

export function OutputMonitor() {
    let params = useParams();

    return <div>{params.clientID}</div>;
}
