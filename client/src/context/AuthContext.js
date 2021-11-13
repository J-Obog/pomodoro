import React, { useState, useEffect, createContext } from 'react';
import axios from 'axios'; 

export const AuthContext = createContext(null); 

const Auth = ({children}) => {
    const [access, setAccessToken] = useState(localStorage.getItem(process.env.REACT_APP_ACCESSTK_KEY));  
    const [refresh, setRefreshToken] = useState(localStorage.getItem(process.env.REACT_APP_REFRESHTK_KEY));
    const [authenticated, setAuthenticated] = useState(false);
    
    useEffect(() => {
        if(!access || !refresh)
            return;  
            
        axios.get(`${process.env.REACT_APP_API_URL}/user`, { headers: { 'Authorization': access }})
        .then(res => {
            setAuthenticated(true); 
        })
        .catch(err => { 
            console.error(err); 
        })

    },[access, refresh])

    const login = (accessToken, refreshToken) => {
        localStorage.setItem(process.env.REACT_APP_ACCESSTK_KEY, accessToken);
        localStorage.setItem(process.env.REACT_APP_REFRESHTK_KEY, refreshToken);
        setAccessToken(accessToken); 
        setRefreshToken(refreshToken);   
        setAuthenticated(true);
    }

    const logout = () => {
        localStorage.setItem(process.env.REACT_APP_ACCESSTK_KEY, null); 
        localStorage.setItem(process.env.REACT_APP_REFRESHTK_KEY, null); 
        setAccessToken(null); 
        setRefreshToken(null);   
        setAuthenticated(false); 
    }

    return (
        <AuthContext.Provider value={{ access, authenticated, logout, login }}>
            {children}
        </AuthContext.Provider>
    )
}; 

export default Auth;