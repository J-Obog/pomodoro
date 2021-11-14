import { React, useState, useContext } from 'react'; 
import axios from 'axios'; 
import { AuthContext } from '../../context/AuthContext';

const Task = (props) => {
    const { access } = useContext(AuthContext);
    const [complete, setCompletion] = useState(props.completed); 
    const [title, setTitle] = useState(props.title);
    const [editing, setEditing] = useState(false); 

    const deleteTask = async () => {
        try {
            const { data } = await axios.delete(`${process.env.REACT_APP_API_URL}/task/${props.id}`, 
            { headers: { 'Authorization': access }});

            props.dispatchDelete(props.id); 
            console.log(data);
        } catch(err) {

        }
    }

    const toggleComplete = async () => {
        try {
            const { data } = await axios.put(`${process.env.REACT_APP_API_URL}/task/${props.id}`, {completed: !complete}, 
            { headers: { 'Authorization': access }});

            setCompletion(!complete); 
            props.dispatchEdit(data); 
            console.log(data);
        } catch(err) {

        }
    }

    return (
        <div className="bg-white rounded-lg p-2 mt-2 border-gray-200 border font-bold text-sm flex items-center justify-between">
           <div>
               <button onClick={toggleComplete}>
                    <span className={`cursor-pointer rounded-md px-1 py-0.5 text-white text-xs ${(complete) ? "bg-green-400" : "bg-red-400"}`}>
                        {(complete) ? "DONE" : "TO DO" }
                    </span>
                </button>
            </div>
            <div>
                <h1 className={`outline-none font-bold ${(complete) ? "line-through" : ""}`}>{title}</h1>
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