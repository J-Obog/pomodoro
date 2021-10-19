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
        <div>
            { tasks.map(task => (
                    <Task key={task.id} {...task}/>
                ))
            }
        </div>
    )
    
}

export default TaskList; 