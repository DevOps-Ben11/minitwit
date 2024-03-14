import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login: React.FC = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault();
        const formData = {
            username: username,
            password: password,
        };
        // TODO : create sim/login
        const response = await fetch('/sim/login', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData),
        });

        if (!response.ok) {
            const errorData = await response.json();
            setError(errorData.error_msg);
        } else {
            setError('');
            navigate("/")
        }
    }

    return (
        <div>
            <h2>Sign in</h2>
            {error && <div className="error"><strong>Error:</strong> {error}</div>}
            <form onSubmit={handleSubmit}>
                <dl>
                    <dt>Username:</dt>
                    <dd><input type="text" name="username" value={username} onChange={(e) => setUsername(e.target.value)} size={30} required /></dd>
                    <dt>Password:</dt>
                    <dd><input type="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} size={30} required /></dd>
                </dl>
                <div className="actions"><input type="submit" value="Sign Up" /></div>
            </form>
        </div>
    );
}

export default Login;
