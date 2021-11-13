import React from 'react';
import { Navigate } from 'react-router-dom';

const PublicRoute = ({ component: Component }) => {
  const isAuth = false; 
   
  return (
      !isAuth ? <Component/> : <Navigate to='/'/>
  )
};


export default PublicRoute;