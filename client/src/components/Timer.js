import React, { useContext } from "react";
import { TimeContext } from "../context/TimeContext";

const Timer = () => {
    const { secs, control } = useContext(TimeContext);

    const setControl = () => {
        control(); 
    }

    return (
        <div className="border-4 rounded-xl bg-white flex-col">
            <div className="text-9xl px-6 py-2">
                <span>{parseInt(secs / 60)}</span>            
                <span>:</span>
                <span>{secs % 60}</span>
            </div>
            <div className="border-t-2">
                <button onClick={setControl} >CONTROL</button>
            </div>
        </div>
    )
}

export default Timer; 