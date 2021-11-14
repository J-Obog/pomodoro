import React from 'react'; 


const TimerMode = (props) => {
    return (
        <div className={`text-md text-center border-b-2 ${(props.active) ? "border-green-500" : "border-gray-200"} w-${props.w} p-4`}>
            <button onClick={() => {props.dispatchModeChange(props.name, props.duration)}}>
                <h1>{props.name.toUpperCase()}</h1>
            </button>
        </div>
    )
}; 


export default TimerMode; 