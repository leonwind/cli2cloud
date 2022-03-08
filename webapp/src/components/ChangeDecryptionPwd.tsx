import {useState} from "react";
import Modal from "react-bootstrap/Modal";
import { Button } from "react-bootstrap";
import { Form } from "react-bootstrap";


export const ChangeDecryptionPwd = ({onSubmit, onClose, show}) => {
    const [password, setPassword] = useState("");

    if (!show) {
        return null;
    }

    const submitAndClose = (password: string) => {
        onSubmit(password);
        setPassword("");
        onClose();
    }

    return (
        <Modal show={show} onHide={onClose}>
            <Modal.Header closeButton>
                <Modal.Title>Enter Password</Modal.Title>
            </Modal.Header>
            
            <Modal.Body>
                <Form>
                    <Form.Group className="mb-3" controlId="formBasicPassword">
                        <Form.Label>Enter your Password to decrypt the output:</Form.Label>
                        <Form.Control type="password" placeholder="Password" 
                            value={password} 
                            onChange={e => setPassword(e.target.value)}/>
                    </Form.Group>
                </Form>
            </Modal.Body>
            
            <Modal.Footer>
                <Button variant="secondary" onClick={onClose}>
                    Close
                </Button>
            
                <Button variant="primary" onClick={() => {submitAndClose(password)}}>
                    Update Password
                </Button>
            </Modal.Footer>
      </Modal>
    )
}