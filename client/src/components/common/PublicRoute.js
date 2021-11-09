import React from 'react';
import { Route, Navigate } from 'react-router-dom';

const PublicRoute = ({ component: Component, restricted, ...rest }) => {
  return (
    <Route {...rest} render={(props) => false && restricted ? <Navigate to='/' /> : <Component {...props} /> }/>
  )
};

export default PublicRoute;