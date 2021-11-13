import axios from 'axios';
import { React, useState, useContext } from 'react'; 
import { Link } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext'; 

const Login = () => {
    const { login } = useContext(AuthContext);
    const [email, setEmail] = useState(''); 
    const [password, setPassword] = useState('');
    //const [error, setError] = useState(null); 

    const attemptLogin = async () => { 
        try {
            const { data } = await axios.post(process.env.REACT_APP_API_URL + '/auth/login', { email: email, password: password });
            login(data.access_token, data.refresh_token); 
        } catch({response}) {
            //handle errors
        }
    }

    return (
        <div className="flex items-center justify-center h-screen">
            <div className="bg-white w-2/6 px-16 py-8 rounded-xl flex flex-col items-center justify-center">
                <div className="mb-16">
                    <h1 className="text-4xl font-bold">Pomodoro</h1>
                </div>
                <div className="mb-10 w-full">
                    <div className="mb-10">
                        <input onInput={e => { setEmail(e.target.value) }} 
                            placeholder="Email" className="w-full outline-none border-b-2"/>
                    </div>
                    <div>
                        <input onInput={e => { setPassword(e.target.value) }} 
                            type="password" placeholder="Password" className="w-full outline-none border-b-2"/>
                    </div>
                </div>
                <div className="text-sm mb-12">
                    <Link to="/register">Not a member yet? Sign up</Link>
                </div>
                <div className="mb-6 w-full">
                    <button onClick={attemptLogin} 
                        className="bg-red-500 py-1 px-3 text-white text-xl rounded-lg w-full">Login</button>
                </div>
            </div>
        </div>
    )
}; 

export default Login; 