import { React, useEffect, useState } from 'react'; 

const Timer = () => {
    const [time, setTime] = useState(25*60); 
    const [running, setRunning] = useState(true); 

    useEffect(() => {
        const invId = setInterval(() => {
            if(running && time > 0) {
                setTime(time - 1); 
            }
        }, 1000); 
        return () => {
            clearInterval(invId);
        }
    }, [running, time, setTime])

    return (
        <div className="bg-white border border-gray-200 w-full px-16 py-8 rounded-xl flex flex-col items-center justify-center"> 
            <div>

            </div>
            <div className="text-gray-800 text-9xl">
                <span>{parseInt(time / 60)}</span>:<span>{(time % 60 < 10) ? "0" : ""}{time % 60}</span>
            </div>
            <div>

            </div>
        </div>
    )
}; 


export default Timer; 