import React, {useState, useEffect} from "react"; 
import axios from "axios"; 
import Task from "./Task"; 
import NewTask from "./NewTask";

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

    const editTask = (data) => {
        const newList = tasks.map(task => { return (task.id == data.id) ? data : task });
        setTasks(newList);  
    }

    const addTask = (data) => {
        const newList = tasks.concat(data);
        setTasks(newList);
    } 

    const deleteTask = (data) => {
        const newList = tasks.filter(task => { return (task.id != data.id) });
        setTasks(newList);
    }

    return (
        <div className="bg-transparent flex flex-col justify-around">
            <div className="bg-black text-white text-lg rounded-lg p-2">
                <h1 className="text-white">Tasks</h1>
            </div>
            <div>
                <NewTask dispatchAddTask={addTask}/>
            </div>
            <div>
                { tasks.map(task => (
                        <Task key={task.id} {...task} 
                            dispatchEditTask={editTask}
                            dispatchDeleteTask={deleteTask}
                        />
                    ))
                }
            </div>
        </div>
    )
    
}

export default TaskList; 