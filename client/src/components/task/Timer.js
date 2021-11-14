import { React, useEffect, useState } from 'react'; 

const Timer = () => {
    const [time, setTime] = useState(1*63); 
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

    const toggleTimer = () => {
        setRunning(!running);
    }

    return (
        <div className="bg-white border border-gray-200 w-full px-0 rounded-xl flex flex-col items-center justify-center"> 
                <div>

                </div>
                <div className={`text-${(time <= 60) ? "red-400": "gray-800"} text-9xl p-10`}>
                    <span>{parseInt(time / 60)}</span>:<span>{(time % 60 < 10) ? "0" : ""}{time % 60}</span>
                </div>
                <div className="h-auto w-full border-t border-gray-200 p-2">
                    <button onClick={toggleTimer} className="bg-red-400 text-white rounded-lg px-12 py-2">
                        <span>{(running) ? "STOP" : "START"}</span>
                    </button>
                </div>
        </div>
    )
}; 


export default Timer; 