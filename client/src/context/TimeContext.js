import React, { createContext, useState, useEffect } from 'react';

export const TimeContext = createContext(null);

const Time = ({ children }) => {
    const [secs, setSeconds] = useState(25 * 60);
    const [running, setRunnng] = useState(true);

    useEffect(() => {
        setInterval(() => {
            if (running) {
                if (secs > 0) setSeconds(secs - 1);
            }
        }, 1000);
    }, [secs, running]);

    const control = () => {
        setRunnng(!running);
    };

    return <TimeContext.Provider value={{ secs, control }}>{children}</TimeContext.Provider>;
};

export default Time;
