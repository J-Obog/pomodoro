import React from 'react';
import { Navigate } from 'react-router-dom';

const PrivateRoute = ({component: Component}) => {
    const isAuth = false; 

    return (
        isAuth ? <Component/> : <Navigate to='/login'/>
    )
}


export default PrivateRoute; 