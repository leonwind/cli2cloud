import styles from "../styles/NavBar.module.css";
import {SearchBar} from "./SearchBar";
import { Navbar } from "react-bootstrap";
import { Nav } from "react-bootstrap";
import { Button, Container } from "react-bootstrap";
import { useState } from "react";
import {ChangeDecryptionPwd} from "./ChangeDecryptionPwd"
import logo from "../assets/cloudWhite.png";

export const NavBar = ({showPasswordBtn, onPasswordSubmit, switchToRawData}) => {
    const [showModal, setShowModal] = useState(false);

    const handleShowModal = () => setShowModal(true);
    const handleCloseModal = () => setShowModal(false);

    return (
        <>
        <Navbar className={styles.body} expand={"sm"} variant={"dark"} collapseOnSelect>
            <Container fluid>
                <Navbar.Brand className={styles.brand} href={"/"}>
                    <img src={logo} alt={"Cli2Cloud"} width={"50"} height={"50"}/> 
                </Navbar.Brand>

                <Navbar.Toggle aria-controls="responsive-navbar-nav"/>
                <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="ms-auto">
                    <SearchBar/>

                    {showPasswordBtn &&
                    <Button variant="dark" className={styles.buttons} onClick={handleShowModal}>
                        Change Password
                    </Button>
                    }

                    <Button variant="dark" className={styles.buttons} onClick={switchToRawData}>
                        Raw
                    </Button>
                </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>

        <ChangeDecryptionPwd show={showModal} onClose={handleCloseModal} onSubmit={onPasswordSubmit}/>
        </>
    )
}