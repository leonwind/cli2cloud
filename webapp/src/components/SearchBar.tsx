import {FormEvent, useState} from "react";
import styles from "../styles/SearchBar.module.css";
import SearchIcon from "@material-ui/icons/Search";
import { useNavigate } from "react-router-dom";
import { Form } from "react-bootstrap";
import { InputGroup } from "react-bootstrap";

export const SearchBar = () => {
    const navigate = useNavigate();
    const [clientID, setClientID] = useState("");

    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        navigate("/" + clientID);
        window.location.reload();
    }

    return (
        <Form onSubmit={handleSubmit}>
                <div className={styles.searchForm}>
                    <Form.Group>
                        <InputGroup>
                            <Form.Control className={styles.searchForm} value={clientID} 
                                onChange={e => setClientID(e.target.value)} type="text" 
                                placeholder="Search your client ID..."/>
                            <InputGroup.Text className={styles.searchForm}>
                                <SearchIcon/>
                            </InputGroup.Text>
                        </InputGroup>
                    </Form.Group>
                </div>
        </Form>
    )
}