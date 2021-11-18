import { React, useContext } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from '../../context/AuthContext';
import { Navbar } from '.';

const PrivateRoute = ({ component: Component }) => {
    const { authenticated } = useContext(AuthContext);

    return authenticated ? (
        <>
            <Navbar />
            <Component />
        </>
    ) : (
        <Navigate to="/login" />
    );
};

export default PrivateRoute;
