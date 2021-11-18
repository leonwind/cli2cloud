import React, {Component} from 'react';

interface State {
    clientID: string
}

export class FrontPage extends Component<{}, State> {

    constructor(props: any) {
        super(props);
        this.state = {
            clientID: ""
        };
    }

    render() {
        return (
            <div>
                Cli2Cloud
            </div>
        )
    }
}