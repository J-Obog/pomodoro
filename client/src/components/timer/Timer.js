import { React, useEffect, useState } from 'react'; 
import TimerMode from './TimerMode';

const timerModes = [
    {
        name: 'Pomodoro', 
        iconColor: 'red',
        duration: 25*60,
    },
    {
        name: 'Short Break', 
        iconColor: 'green',
        duration: 5*60,
    },
    {
        name: 'Long Break', 
        iconColor: 'blue',
        duration: 15*60,
    }
]

const Timer = () => {
    const [time, setTime] = useState(1*63); 
    const [running, setRunning] = useState(false); 
    const [currentMode, setCurrentMode] = useState('Pomodoro'); 

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

    const resetTimer = () => {
        setRunning(false);
        setTime(25*60); 
    }

    const handleModeChange = (newMode, newDuration) => { 
        setRunning(false);
        setCurrentMode(newMode);
        setTime(newDuration);
    }

    return (
        <div className="bg-white border border-gray-200 w-3/4 px-0 rounded-xl flex flex-col items-center justify-center"> 
                <div className="w-full flex flex-row">
                    { timerModes.map(mode => (
                            <TimerMode key={mode.name} {...mode}
                                w={"1/3"}
                                active={(mode.name === currentMode) ? true : false}
                                dispatchModeChange={handleModeChange}
                            />
                        ))
                    }
                </div>
                <div className={`text-${(time <= 60) ? "red-400": "gray-800"} text-8xl p-10`}>
                    <span>{parseInt(time / 60)}</span>:<span>{(time % 60 < 10) ? "0" : ""}{time % 60}</span>
                </div>
                <div className="h-auto w-full border-t border-gray-200 py-3">
                    <button onClick={toggleTimer} className="bg-red-400 text-white rounded-lg px-12 py-2 mx-6">
                        <span>{(running) ? "STOP" : "START"}</span>
                    </button>
                    <button onClick={resetTimer} className="bg-white text-red-400 border border-gray-200 rounded-lg px-12 py-2">
                        RESET
                    </button>
                </div>
        </div>
    )
}; 


export default Timer; 