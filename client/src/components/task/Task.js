import { React, useState, useContext } from 'react';
import axios from 'axios';
import { AuthContext } from '../../context/AuthContext';

const Task = (props) => {
    const { access } = useContext(AuthContext);
    const [complete, setCompletion] = useState(props.completed);
    const [title, setTitle] = useState(props.title);
    //const [editing, setEditing] = useState(false);

    const deleteTask = async () => {
        try {
            const { data } = await axios.delete(`${process.env.REACT_APP_API_URL}/task/${props.id}`, { headers: { Authorization: access } });

            props.dispatchDelete(props.id);
            console.log(data);
        } catch (err) {}
    };

    const toggleComplete = async () => {
        try {
            const { data } = await axios.put(
                `${process.env.REACT_APP_API_URL}/task/${props.id}`,
                { completed: !complete },
                { headers: { Authorization: access } },
            );

            setCompletion(!complete);
            props.dispatchEdit(data);
            console.log(data);
        } catch (err) {}
    };

    return (
        <div className="bg-white rounded-lg p-2 mt-2 border-gray-200 border text-sm flex items-center justify-between">
            <div>
                <button onClick={toggleComplete}>
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className={`h-7 w-7 fill-current text-green-500 opacity-${complete ? '100' : '25'}`}
                        viewBox="0 0 20 20">
                        <path
                            fillRule="evenodd"
                            d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                            clipRule="evenodd"
                        />
                    </svg>
                </button>
            </div>
            <div>
                <h1 className={`outline-none ${complete ? 'line-through' : ''}`}>{title}</h1>
            </div>
            <div>
                <button onClick={deleteTask}>
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="#4B5563">
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                        />
                    </svg>
                </button>
            </div>
        </div>
    );
};

export default Task;
