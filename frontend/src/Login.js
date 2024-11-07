import React, { useState } from "react";
import api from './api';

const Login = ({ setUser }) => {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [role, setRole] = useState('')

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            // Call the backend login endpoint
            setRole('super_admin');
            const response = await api.post('/login', { username, password, role });
            
            // Update the user state with the response data
            setUser(response.data);
            alert('Login Successful!');
        } catch (error) {
            console.error("Login failed:", error);
            alert('Login failed!');
        }
    };

    return (
        <form onSubmit={handleLogin}>
            <h2>Login</h2>
            <div>
                <label>Username:</label>
                <input type="text" value={username} placeholder="Enter your username" onChange={(e) => setUsername(e.target.value)} />
            </div>
            <div>
                <label>Password:</label>
                <input type="text" value={password} placeholder="Enter your password" onChange={(e) => setPassword(e.target.value)} />
            </div>
            <button type="submit">Login</button>
        </form>
    );
};

export default Login;