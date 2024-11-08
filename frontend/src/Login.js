import React, { useState } from "react";
import api from './api';
import "./Login.css"

import { InputText } from 'primereact/inputtext';
import { FloatLabel } from 'primereact/floatlabel';
import { Button } from 'primereact/button';

const Login = ({ setUser }) => {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            // Call the backend login endpoint
            const response = await api.post('/login', { username, password });
            
            // Update the user state with the response data
            setUser(response.data);
            // alert('Login Successful!');
        } catch (error) {
            console.error("Login failed:", error);

            // Extract detailed error information
            const statusCode = error.response?.status;
            const responseData = error.response?.data;

            // Show a more detailed error alert
            alert(`Login failed:
            Error Message: ${error}
            Status Code: ${statusCode}
            Response Data: ${JSON.stringify(responseData)}`);
        }
    };

    // Prevent space in input box
    const onKeyDown = (event) => {
        if (event.code === 'Space') event.preventDefault()
    }

    return (
        <div className="login-page-wrap">
            <div className="title">
                <h1>Event Attendance System</h1>
            </div>
            <div className="subtitle">
                <h2>Login</h2>
            </div>
            

            <form onSubmit={handleLogin} className="form-wrap">
                <div>
                    <FloatLabel className="input-wrap">
                        <InputText type="text" value={username} 
                                placeholder="Enter your username"
                                className="input-box"
                                onChange={(e) => setUsername(e.target.value)}
                                onKeyDown={onKeyDown}
                                autoCorrect="off" 
                                spellCheck="false" 
                                autoComplete="username"
                                autoCapitalize="off" />
                        <label htmlFor="username">Username</label>
                    </FloatLabel>  
                </div>
                <div>
                    <FloatLabel className="input-wrap">
                        <InputText type="text" value={password} 
                                placeholder="Enter your password" 
                                className="input-box"
                                onChange={(e) => setPassword(e.target.value)}
                                onKeyDown={onKeyDown}
                                autoCorrect="off" 
                                spellCheck="false" 
                                autoComplete="username"
                                autoCapitalize="off" />
                        <label htmlFor="password">Password</label>
                    </FloatLabel>
                </div>
                <Button type="submit" rounded className="login-btn">Log in</Button>
            </form>
        </div>
    );
};

export default Login;