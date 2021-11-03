import axios from "axios";
import React, { useState } from "react"; 

const NewTask = (props) => {
    const [title, setTitle] = useState("");

    const addTask = () => {
        axios.post("http://localhost:8000/api/task", { title: title })
        .then(({data}) => {
            props.dispatchAddTask(data); 
            setTitle("");
        })  
        .catch(err => {
            console.error(err);
        })
    }

    const changeTitle = ({target}) => {
        setTitle(target.value);
    }
    
    return (
        <div className="bg-gray-100 rounded-lg p-2 mt-2 font-bold text-sm flex items-center justify-between">
            <div>
                <input onInput={changeTitle} value={title} placeholder="New Task" className="font-bold bg-transparent outline-none"/>
            </div>
            <div>
                <button onClick={addTask}>
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="#0EA5E9">
                        <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V7z" clipRule="evenodd" />
                    </svg>
                </button>
            </div>
        </div>
    )
}

export default NewTask; 