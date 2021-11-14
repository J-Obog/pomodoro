import { React } from 'react'; 
import TaskList from '../components/task/TaskList'; 
import Timer from '../components/task/Timer'; 


const Dashboard = () => {    
    return (
        <div className="timer-app flex h-screen">
            <div className="flex flex-col items-center w-4/6 m-3">
                <Timer/>
            </div>
            <div className="flex flex-col w-2/6 m-3">
                <TaskList/>
            </div>
        </div>
    )
}; 

export default Dashboard; 