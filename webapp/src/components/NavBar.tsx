import styles from "../styles/NavBar.module.css";
import {SearchBar} from "./SearchBar";
import { Navbar } from "react-bootstrap";
import { Nav } from "react-bootstrap";
import { Button } from "react-bootstrap";
import { useState } from "react";
import {ChangeDecryptionPwd} from "./ChangeDecryptionPwd"
import logo from "../assets/cloudWhite.png";

export const NavBar = ({showPasswordBtn, onPasswordSubmit, switchToRawData}) => {
    const [showModal, setShowModal] = useState(false);

    const handleShowModal = () => setShowModal(true);
    const handleCloseModal = () => setShowModal(false);

    return (
        <>
        <Navbar className={styles.body} expand={"md"} variant={"dark"} collapseOnSelect>
            <Navbar.Brand className={styles.brand} href={"/"}>
                <img src={logo} alt={"Logo"} width={"50"} height={"50"}/>
            </Navbar.Brand>

            <Navbar.Toggle aria-controls="responsive-navbar-nav"/>
            <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="ms-auto">
                <SearchBar/>

                <Button variant="dark" className={styles.buttons} onClick={switchToRawData}>
                    Raw
                </Button>

                {showPasswordBtn &&
                <Button variant="dark" className={styles.buttons} onClick={handleShowModal}>
                    Change Password
                </Button>
                }

                
            </Nav>
            </Navbar.Collapse>
            
        </Navbar>

        <ChangeDecryptionPwd show={showModal} onClose={handleCloseModal} onSubmit={onPasswordSubmit}/>
        </>
    )
}