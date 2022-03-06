import styles from "../styles/NavBar.module.css";
import {SearchBar} from "./SearchBar";
import { Navbar } from "react-bootstrap";
import { Nav } from "react-bootstrap";
import { Button } from "react-bootstrap";
import { useState } from "react";
import {ChangeDecryptionPwd} from "./ChangeDecryptionPwd"

export const NavBar = ({showPasswordBtn, onPasswordSubmit}) => {
    const [showModal, setShowModal] = useState(false);

    const handleShowModal = () => setShowModal(true);
    const handleCloseModal = () => setShowModal(false);

    return (
        <>
        <Navbar className={styles.body} expand={"md"} variant={"dark"}>
            <Navbar.Brand className={styles.brand} href={"/"}>
                Cli2Cloud
            </Navbar.Brand>

            <Nav>
                <SearchBar/>
            </Nav>

            {showPasswordBtn &&
            <Nav className="ms-auto">
                <Button onClick={handleShowModal}>
                    Change Password
                </Button>
            </Nav>
            }
        </Navbar>

        <ChangeDecryptionPwd show={showModal} onClose={handleCloseModal} onSubmit={onPasswordSubmit}/>
        </>
    )
}