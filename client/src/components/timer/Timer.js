import { React, useEffect, useState } from 'react'; 
import TimerMode from './TimerMode';

const timerModes = [
    {
        name: 'Pomodoro', 
        theme: 'red',
        duration: 25*60,
    },
    {
        name: 'Short Break', 
        theme: 'green',
        duration: 5*60,
    },
    {
        name: 'Long Break', 
        theme: 'blue',
        duration: 15*60,
    }
]

const Timer = () => {
    const [time, setTime] = useState(25*60); 
    const [running, setRunning] = useState(false); 
    const [currentMode, setCurrentMode] = useState('Pomodoro'); 

    useEffect(() => {
        const invId = setInterval(() => {
            if(running && time > 0) {
                setTime(time - 1); 
            }

            if(time === 0) { 
                setRunning(false);
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
        const t = timerModes.filter(m => { return m.name === currentMode }); 
        setRunning(false);
        setTime(t[0].duration); 
    }

    const handleModeChange = (newMode, newDuration) => { 
        if(running) {
            const isConfirmed = window.confirm("The timer is still running, are you sure you want to move onto the next round?");

            if(!isConfirmed)
                return; 
        }

        setRunning(false);
        setCurrentMode(newMode);
        setTime(newDuration);
    }

    return (
        <div className="bg-white border border-gray-200 w-3/5 px-0 rounded-xl flex flex-col items-center justify-center"> 
                <div className="w-full flex flex-row">
                    { timerModes.map(mode => (
                            <TimerMode key={mode.name} {...mode}
                                w={`1/${timerModes.length}`}
                                active={(mode.name === currentMode) ? true : false}
                                dispatchModeChange={handleModeChange}
                            />
                        ))
                    }
                </div>
                <div className={`bg-gray-100 w-full text-center text-${(time <= 60) ? "red-400": "gray-600"} text-7xl p-10`}>
                    <span>{parseInt(time / 60)}</span>|<span>{(time % 60 < 10) ? "0" : ""}{time % 60}</span>
                </div>
                <div className="text-xs h-auto w-full border-t border-gray-200 py-3">
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