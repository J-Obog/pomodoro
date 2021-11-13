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
            await axios.post(process.env.REACT_APP_API_URL + '/auth/register', { email: email, password: password});
            navigate('/login');
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
                        <input onInput={e => {setEmail(e.target.value)}}
                            placeholder="Email" className="w-full outline-none border-b-2"/>
                    </div>
                    <div>
                        <input onInput={e => {setPassword(e.target.value)}}
                            type="password" placeholder="Password" className="w-full outline-none border-b-2"/>
                    </div>
                </div>
                <div className="text-sm mb-12">
                    <Link to="/login">Already have an account? Log In</Link>
                </div>
                <div className="mb-6 w-full">
                    <button onClick={attemptRegistration} 
                    className="bg-red-500 py-1 px-3 text-white text-xl rounded-lg w-full">Sign Up</button>
                </div>
            </div>
        </div>
    )
}; 

export default Register; 