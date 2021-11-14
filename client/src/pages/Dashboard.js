import { React } from 'react'; 
import TaskList from '../components/task/TaskList'; 


const Dashboard = () => {    
    return (
        <div className="timer-app flex h-screen">
            <div className="w-4/6 m-3">
                {/* Timer */}
            </div>
            <div className="w-2/6 m-3">
                <TaskList/>
            </div>
        </div>
    )
}; 

export default Dashboard; 