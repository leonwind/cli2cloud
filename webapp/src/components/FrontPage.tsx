import React, {ChangeEvent, Component, FormEvent} from "react";
import styles from "../styles/FrontPage.module.css";
import 'bootstrap/dist/css/bootstrap.min.css';
import { Form } from "react-bootstrap";
import { Navigate } from "react-router-dom";

interface State {
    clientID: string,
    redirect: boolean
}

export class FrontPage extends Component<{}, State> {
    
    constructor(props: any) {
        super(props);
        this.state = {
            clientID: "",
            redirect: false
        };

        this.handleClientIDChange = this.handleClientIDChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    private handleClientIDChange(event: ChangeEvent<HTMLInputElement>) {
        this.setState({clientID: event.target.value});
    }

    private handleSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();
        console.log(this.state.clientID);
        this.setState({redirect: true})
    } 

    render() {
        if (this.state.redirect) {
            return (
                <Navigate to={"/" + this.state.clientID}/>
            )
        }
        return (
            <div className={styles.body}>

                <h1 className={styles.headline}>
                    Cli2Cloud
                </h1>

                <br/>
                <br/>

                <Form onSubmit={this.handleSubmit}>
                <div className={styles.searchForm}> 
                        <Form.Control className={styles.searchForm} value={this.state.clientID} onChange={this.handleClientIDChange} size="lg" type="text" 
                            placeholder="Enter your terminal ID..."/>
                </div>
                </Form>

                <br/>
                <br/>

                <h6 className={styles.mediumText}>
                    Monitor and Share Your Yerminal Output with Everyone from Everywhere in Realtime.
                </h6>
            </div>
        )
    }
}