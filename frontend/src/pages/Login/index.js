import React from 'react'
import { Form, Button, Container } from "react-bootstrap";
import { Link } from "react-router-dom";
import "./style.css"

export const Login = () => {
    return (
        <Container>
            <div className='d-flex justify-content-center align-items-center login'>

                <Form className="loginform" >
                    <h3 className="masuk">Masuk</h3>
                    <Form.Group className="formgroup">
                        <Form.Label>Email</Form.Label>
                        <Form.Control id="email" className="form-input" type="email" required />
                    </Form.Group>

                    <Form.Group className="formgroup">
                        <Form.Label>Kata Sandi</Form.Label>
                        <Form.Control id="password" type="password" required />
                    </Form.Group>
                    <Button className="loginbutton w-100" type="submit">
                        MASUK
                    </Button>
                </Form>

            </div>
        </Container>
    )
}

export default Login;
