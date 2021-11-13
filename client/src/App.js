import React from 'react';

import { Routes, Route, BrowserRouter} from 'react-router-dom'; 
import { PublicRoute, PrivateRoute } from './components/common';

import { Dashboard, Login, Register } from './pages'; 

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path='/login' element={<PublicRoute component={Login}/>}/>
        <Route exact path='/register' element={<PublicRoute component={Register}/>}/>
        <Route exact path='/' element={<PrivateRoute component={Dashboard}/>}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
