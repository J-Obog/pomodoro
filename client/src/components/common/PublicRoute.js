import { React, useContext } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from '../../context/AuthContext';

const PublicRoute = ({ component: Component }) => {
  const { authenticated } = useContext(AuthContext);
   
  return (
      !authenticated ? <Component/> : <Navigate to='/'/>
  )
};


export default PublicRoute;