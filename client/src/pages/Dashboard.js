import { React } from 'react'; 
import { TaskList } from '../components/Task'; 
import { Timer } from '../components/Timer'; 


const Dashboard = () => {    
    return (
        <div className="flex h-screen p-4">
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