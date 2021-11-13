import { React, useState } from 'react'; 
import Task from './Task'; 

const TaskList = (props) => {
    const [tasks, setTasks] = useState(props.userTasks);

    return (
        <div className="bg-transparent flex flex-col justify-around">
            <div className="bg-black text-white text-lg rounded-lg p-2 mb-8">
                <h1>Tasks</h1>
            </div>
            <div className="mb-10">
                <input placeholder="New Task" className="w-full outline-none border-b-2"/>
            </div>
            <div>
                { tasks.map(task => ( <Task key={task.id} {...task}/> )) }
            </div>
        </div>
    )
    
}

export default TaskList; 