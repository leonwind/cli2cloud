import React, {ChangeEvent, Component, FormEvent} from "react";
import styles from "../styles/FrontPage.module.css";
import 'bootstrap/dist/css/bootstrap.min.css';
import { Form } from "react-bootstrap";
import SearchIcon from "@material-ui/icons/Search"

interface State {
    clientID: string
}

export class FrontPage extends Component<{}, State> {

    constructor(props: any) {
        super(props);
        this.state = {
            clientID: ""
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
    } 

    render() {
        return (
            <div className={styles.body}>
                <h1 className={styles.headline}>
                    Cli2Cloud
                </h1>
                
                <div className={styles.searchForm}>
                <Form onSubmit={this.handleSubmit}>
                    <Form.Group className="mb-3">
                        <SearchIcon fontSize="large" className={styles.searchIcon}/>
                        <Form.Control value={this.state.clientID} onChange={this.handleClientIDChange} size="lg" type="text" 
                            placeholder="Enter your terminal ID" />
                        <br/>
                        <h6 className={styles.mediumText}>
                            Monitor your terminal output from everywhere.
                        </h6>
                    </Form.Group>
                </Form>
                </div>
            </div>
        )
    }
}