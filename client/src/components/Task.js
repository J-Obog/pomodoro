import React, { useState } from "react"; 
import axios from "axios"; 

const Task = (props) => {
    const [complete, setCompletion] = useState(props.completed); 
    const [title, setTitle] = useState(props.title);

    const toggle = () => {
        axios.put(`http://localhost:8000/api/tasks/${props.id}`, { title: title, completed: !complete })
        .then(({data}) => {
            setCompletion(!complete);
        })
        .catch(err => {
            console.error(err); 
        }) 
    }   

    return (
        <div className="bg-white rounded-lg p-2 mt-2 font-bold text-sm flex items-center justify-between">
           <div>
               <span onClick={toggle} className={`rounded-md px-1 py-0.5 text-white text-xs ${(complete) ? "bg-green-400" : "bg-red-400"}`}>
                   {(complete) ? "DONE" : "TO DO" }
                </span>
            </div>
            <div>
                <h1 className={`font-bold ${(complete) ? "line-through" : ""}`}>{title}</h1>
            </div>
            <div>

            </div>
        </div>
    )
}

export default Task; 