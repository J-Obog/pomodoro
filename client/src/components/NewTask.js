import axios from "axios";
import React, { useState } from "react"; 

const NewTask = () => {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState(""); 

    const addNewTask = () => {
        axios.post("http://localhost:8000/api/tasks", {title, description})
        .then(({data}) => {
            console.log(data); 
        })  
        .catch(err => {
            console.error(err);
        })
    }

    const changeTitle = ({target}) => {
        setTitle(target.value);
    }

    const changeDescription = ({target}) => {
        setDescription(target.value);
    }
    
    return (
        <div className="bg-gray-100 text-white rounded-lg p-2 mt-3">
            <div className="text-gray-400 text-center text-sm">
                <input onInput={changeTitle} type="text" placeholder="Title" className="text-center font-bold bg-transparent outline-none"/>
                <input onInput={changeDescription} type="text" placeholder="Description" className="text-center bg-transparent outline-none"/>
            </div>
            <div className="flex items-end">
                <button onClick={addNewTask} className="bg-gray-900 rounded-xl text-xs py-0.5 px-1">ADD</button>
            </div>
        </div>
    )
}

export default NewTask; 