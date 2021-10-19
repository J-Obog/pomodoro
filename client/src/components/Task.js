import React from "react"; 

const Task = ({id, title, description, completed}) => {
    return (
        <div className="bg-white text-white rounded-lg p-1 mt-3">
            <div className="text-black text-center text-sm">
                <h1 className="font-bold">{title}</h1>
                <p>{description}</p>
            </div>
        </div>
    )
}

export default Task; 