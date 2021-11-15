import React from 'react'; 


const TimerMode = (props) => {
    return (
        <div className={`transition duration-500 text-md text-center border-b-2 ${(props.active) ? "border-green-500" : "border-gray-200"} w-${props.w} p-4`}>
            <div className="flex justify-center items-center">
                <svg xmlns="http://www.w3.org/2000/svg" width="15" height="15" fill={props.theme}>
                    <path d="M14 7.5a6.5 6.5 0 11-13 0 6.5 6.5 0 0113 0z"></path>
                </svg>
                <button className="ml-3" onClick={() => {props.dispatchModeChange(props.name, props.duration)}}>
                    <h1>{props.name.toUpperCase()}</h1>
                </button>
            </div>
        </div>
    )
}; 


export default TimerMode; 