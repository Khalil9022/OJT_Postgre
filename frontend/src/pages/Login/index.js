import axios from 'axios';
import React, { useState } from 'react'
import { Form, Button } from "react-bootstrap";
import "./style.css"
import swal from "sweetalert";

export const Login = () => {

    const [userDetails, setUserDetails] = useState({
        email: "",
        password: "",
    });

    const handleChange = (event) => {
        setUserDetails({ ...userDetails, [event.target.id]: event.target.value });
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        axios
            .post("https://reqres.in/api/login", userDetails)
            .then((res) => {
                localStorage.setItem("isLoggedin", true);
                localStorage.setItem("token", res.data.token);
                swal("Sukses", "Berhasil masuk", "success").then(
                    window.location.reload()
                );

            })
            .catch((error) => {
                swal("Gagal login", "Id / Password Salah", "warning");
            });
    };

    return (
        <div className='d-flex justify-content-center align-items-center login'>

            <Form className="loginform" onSubmit={handleSubmit} >
                <h3 className="masuk">Masuk</h3>
                <Form.Group className="formgroup">
                    <Form.Label>Email</Form.Label>
                    <Form.Control id="email" className="form-input" type="email" required onChange={handleChange} />
                </Form.Group>

                <Form.Group className="formgroup">
                    <Form.Label>Kata Sandi</Form.Label>
                    <Form.Control id="password" type="password" required onChange={handleChange} />
                </Form.Group>
                <Button className="loginbutton w-100" type="submit">
                    MASUK
                </Button>
            </Form>

        </div>
    )
}

export default Login;
