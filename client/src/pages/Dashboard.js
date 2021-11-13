import { React, useEffect, useContext, useState } from 'react'; 
import axios from 'axios';
import TaskList from '../components/task/TaskList'; 
import { AuthContext } from '../context/AuthContext';


const Dashboard = () => {
    const { access } = useContext(AuthContext); 
    const [tasks, setTasks] = useState([]); 

    useEffect(() => {
        axios.get(`${process.env.REACT_APP_API_URL}/task`, { headers: { 'Authorization': access }})
        .then(({data}) => {
            setTasks(data); 
        })
        .catch(err => {})
    }, [access])
    
    return (
        <div className="timer-app flex h-screen">
            <div className="w-4/6 m-3">
                {/* Timer */}
            </div>
            <div className="w-2/6 m-3">
                <TaskList userTasks={tasks}/>
            </div>
        </div>
    )
}; 

export default Dashboard; 