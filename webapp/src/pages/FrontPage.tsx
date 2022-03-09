import React, {ChangeEvent, Component, FormEvent} from "react";
import styles from "../styles/FrontPage.module.css";
import 'bootstrap/dist/css/bootstrap.min.css';
import { Form } from "react-bootstrap";
import { Navigate } from "react-router-dom";
import logo from "../assets/cloudWhite.png";
import SearchIcon from "@material-ui/icons/Search";
import { InputGroup } from "react-bootstrap";
import {Documentation} from "./Documentation"

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
                <h1 className={styles.logo}>
                    <img src={logo} alt={"Logo"} width={"100"} height={"100"}/>
                </h1>
                
                <h1 className={styles.headline}>
                    Cli2Cloud
                </h1>

                <br/>
                <br/>

                <Form onSubmit={this.handleSubmit}>
                <div className={styles.searchForm}> 

                    <Form.Group>
                        <InputGroup>
                        <Form.Control className={styles.searchForm} value={this.state.clientID} 
                            onChange={this.handleClientIDChange} size="lg" type="text" 
                            placeholder="Enter your client ID..."/>
                        <InputGroup.Text className={styles.searchLogo}>
                                <SearchIcon/>
                            </InputGroup.Text>
                        </InputGroup>
                    </Form.Group>
                </div>
                </Form>

                <br/>
                <br/>

                <h6 className={styles.mediumText}>
                    Monitor and Share Your Terminal Output with Everyone from Everywhere in Realtime.
                </h6>

                <br/>
                <br/>
                <br/>

                <Documentation/>
            </div>
        )
    }
}