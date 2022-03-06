import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import styles from "../styles/NavBar.module.css";
import { Form } from "react-bootstrap";


export const NavBar = () => {

    return (
        <>
            <Navbar className={styles.body} expand={"md"} variant={"dark"}>
                <Navbar.Brand className={styles.brand} href={"/"}>
                    Cli2Cloud
                </Navbar.Brand>

                <Nav>
                    <Form>
                        <div className={styles.searchForm}> 
                            <Form.Control className={styles.searchForm}
                                type="text" 
                                placeholder={"Search new ID..."}/>
                        </div>
                    </Form>
                </Nav>
            </Navbar>
        </>
    )
}