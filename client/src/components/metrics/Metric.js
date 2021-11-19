import React from 'react';

const Metric = (props) => {
    return (
        <div className="flex flex-row font-bold items-center justify-between bg-white px-5 py-3 my-3 shadow-md rounded-md">
            <div className="text-2xl">
                <h1>{props.metric}</h1>
            </div>
            <div className="text-4xl text-green-500">
                <h1>{props.value}</h1>
            </div>
        </div>
    );
};

export default Metric;
