import React from "react"; 

const Task = ({id, title, completed}) => {
    return (
        <div className="bg-white rounded-lg p-2 mt-2 text-sm flex items-center justify-between">
            <input type="checkbox" className="cursor-pointer"/>
            <h1 className="font-bold">{title}</h1>
            <button className="bg-red-500 text-white rounded-full h-5 w-5">x</button>
        </div>
    )
}

export default Task; 