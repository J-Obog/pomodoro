import React, { useState } from "react"; 
import axios from "axios"; 

const Task = (props) => {
    const [complete, setCompletion] = useState(props.completed); 
    const [title, setTitle] = useState(props.title);

    const toggle = () => {
        axios.put(`http://localhost:8000/api/tasks/${props.id}`, { title: title, completed: !complete })
        .then(({data}) => {
            setCompletion(!complete);
            props.dispatchEditTask(data);
        })
        .catch(err => {
            console.error(err); 
        }) 
    }   

    const deleteTask = () => {
        axios.delete(`http://localhost:8000/api/tasks/${props.id}`)
        .then(({data}) => {
            props.dispatchDeleteTask(data); 
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
                <button onClick={deleteTask}>
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="#F43F5E">
                        <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                    </svg>
                </button>
            </div>
        </div>
    )
}

export default Task; 