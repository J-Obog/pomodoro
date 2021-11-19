import axios from 'axios';
import { React, useState, useContext } from 'react';
import { Link } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';

const Login = () => {
    const { login } = useContext(AuthContext);
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(null);

    const attemptLogin = async () => {
        try {
            const { data } = await axios.post(process.env.REACT_APP_API_URL + '/auth/login', { email: email, password: password });
            login(data.access_token, data.refresh_token);
        } catch ({ response }) {
            setError(response.data);
        }
    };

    const inputKeyPressed = (e) => {
        if (e.key === 'Enter') {
            attemptLogin();
        }
    };

    return (
        <div className="flex items-center justify-center w-full">
            <div className="auth-form-modal">
                <div className="mb-16">
                    <h1 className="text-4xl font-bold">Pomodoro</h1>
                </div>
                <div className="mb-10 w-full">
                    <div className="mb-10">
                        <input
                            onKeyPress={inputKeyPressed}
                            onInput={(e) => {
                                setEmail(e.target.value);
                            }}
                            placeholder="Email"
                            className="auth-form-input"
                        />

                        <div className="auth-form-input-error">
                            <span>{error?.email || ''}</span>
                        </div>
                    </div>
                    <div>
                        <input
                            onKeyPress={inputKeyPressed}
                            onInput={(e) => {
                                setPassword(e.target.value);
                            }}
                            type="password"
                            placeholder="Password"
                            className="auth-form-input"
                        />

                        <div className="auth-form-input-error">
                            <span>{error?.password || ''}</span>
                        </div>
                    </div>
                </div>
                <div className="text-sm mb-12">
                    <Link to="/register">
                        Not a member yet? <b>Sign up</b>
                    </Link>
                </div>
                <div className="mb-6 w-full">
                    <button onClick={attemptLogin} className="auth-form-submit-btn">
                        Login
                    </button>
                </div>
            </div>
        </div>
    );
};

export default Login;
