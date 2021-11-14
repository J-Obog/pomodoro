import React from 'react'; 


const TimerMode = (props) => {
    return (
        <div className={`text-lg text-center border-b-2 ${(props.active) ? "border-blue-500" : "border-gray-200"} w-${props.w} p-4`}>
            <button>
                <h1>{props.name.toUpperCase()}</h1>
            </button>
        </div>
    )
}; 


export default TimerMode; 