import React, { useState } from 'react';
import "../style.css";
import { useNavigate } from 'react-router-dom';

/*
    This component for the registration handle few things for now:
        - Check if both passwords are identical and send an error if not.
        - Send a registration form to the backend via /sim/register (for now probably to change).
        - Is informed is the user has been registered and send to timeline or an error message.

    Issues I have:
        - Is the JSON form sent to backend secured ?
        - Is the user sign in after the registration or should we redirect to login page? 

    Except for these issues I thing that this component do the job for now.
*/

const Register: React.FC = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [password2, setPassword2] = useState('');
    const [error, setError] = useState('');

    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault();
        const formData = {
            username: username,
            email: email,
            password: password,
        };

        if (password !== password2) {
            setError("Passwords do not match.");
            return;
        }

        const response = await fetch('/sim/register', {
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
            <h2>Sign Up</h2>
            {error && <div className="error"><strong>Error:</strong> {error}</div>}
            <form onSubmit={handleSubmit}>
                <dl>
                    <dt>Username:</dt>
                    <dd><input type="text" name="username" value={username} onChange={(e) => setUsername(e.target.value)} size={30} required /></dd>
                    <dt>E-Mail:</dt>
                    <dd><input type="text" name="email" value={email} onChange={(e) => setEmail(e.target.value)} size={30} required /></dd>
                    <dt>Password:</dt>
                    <dd><input type="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} size={30} required /></dd>
                    <dt>Password <small>(repeat)</small>:</dt>
                    <dd><input type="password" name="password2" value={password2} onChange={(e) => setPassword2(e.target.value)} size={30} required /></dd>
                </dl>
                <div className="actions"><input type="submit" value="Sign Up" /></div>
            </form>
        </div>
    );
}

export default Register;
