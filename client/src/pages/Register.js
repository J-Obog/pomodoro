import { React, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';

const Register = () => {
    const navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    //const [error, setError] = useState(null);

    const attemptRegistration = async () => {
        try {
            await axios.post(process.env.REACT_APP_API_URL + '/auth/register', {
                email: email,
                password: password,
            });
            navigate('/login');
        } catch ({ response }) {
            //handle errors
        }
    };

    return (
        <div className="flex h-full justify-center items-center">
            <div className="auth-form-modal">
                <div className="mb-16">
                    <h1 className="text-4xl font-bold">Pomodoro</h1>
                </div>
                <div className="mb-10 w-full">
                    <div className="mb-10">
                        <input
                            onInput={(e) => {
                                setEmail(e.target.value);
                            }}
                            placeholder="Email"
                            className="auth-form-input"
                        />
                    </div>
                    <div>
                        <input
                            onInput={(e) => {
                                setPassword(e.target.value);
                            }}
                            type="password"
                            placeholder="Password"
                            className="auth-form-input"
                        />
                    </div>
                </div>
                <div className="text-sm mb-12">
                    <Link to="/login">
                        Already have an account? <b>Log In</b>
                    </Link>
                </div>
                <div className="mb-6 w-full">
                    <button onClick={attemptRegistration} className="auth-form-submit-btn">
                        Sign Up
                    </button>
                </div>
            </div>
        </div>
    );
};

export default Register;
