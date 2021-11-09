import React from 'react';

//routing
import { Routes, BrowserRouter as Router } from 'react-router-dom'; 
import { PublicRoute, PrivateRoute } from './components/common';

function App() {
  return (
    <Router>
      <Routes>
        <PublicRoute path='/login' component={<><div></div></>} />
      </Routes>
    </Router>
  );
}

export default App;
