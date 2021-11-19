import { React } from 'react';
import { TaskList } from '../components/Task';
import { Timer } from '../components/Timer';

const Dashboard = () => {
    return (
        <div className="flex flex-row items-start py-10 px-10">
            <div className="w-4/6">
                <Timer />
            </div>
            <div className="w-2/6">
                <TaskList />
            </div>
        </div>
    );
};

export default Dashboard;
