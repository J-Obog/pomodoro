import React, {useState, useEffect} from "react"; 
import axios from "axios"; 
import Task from "./Task"; 


const TaskList = () => {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:8000/api/tasks")
        .then(({data}) => {
            setTasks(data);  
        })
        .catch(err => {
            console.error(err); 
        })
    }, [])

    return (
        <div className="bg-transparent flex flex-col justify-around">
            <div className="bg-black text-white text-lg rounded-lg p-2">
                <h1 className="text-white">Tasks</h1>
            </div>
            <div>
                { tasks.map(task => (
                        <Task key={task.id} {...task}/>
                    ))
                }
            </div>
        </div>
    )
    
}

export default TaskList; 