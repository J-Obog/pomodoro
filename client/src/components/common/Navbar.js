import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => {
    return (
        <div className="bg-red-400 py-5 px-4 flex flex-row items-start justify-space-between">
            <div className="mx-6 text-white">
                <Link to="/">Pomodoro</Link>
            </div>
            <div className="mx-6 text-white">
                <Link to="/metrics">Metrics</Link>
            </div>
        </div>
    );
};

export default Navbar;
