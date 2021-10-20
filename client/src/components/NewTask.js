import axios from "axios";
import React, { useState } from "react"; 

const NewTask = ({dispatch}) => {
    const [title, setTitle] = useState("");

    const addNewTask = () => {
        axios.post("http://localhost:8000/api/tasks", { title })
        .then(({data}) => {
            dispatch(data); 
        })  
        .catch(err => {
            console.error(err);
        })
    }

    const changeTitle = ({target}) => {
        setTitle(target.value);
    }
    
    return (
        <div className="bg-gray-100 text-white rounded-lg p-2 mt-3">
            <div className="text-gray-400 text-center text-sm flex justify-between">
                <input onInput={changeTitle} type="text" placeholder="New Task" className="font-bold bg-transparent outline-none"/>
                <button onClick={addNewTask} className="bg-blue-500 text-white font-bold rounded-full h-5 w-5">+</button>
            </div>
        </div>
    )
}

export default NewTask; 