import React from 'react';
import { MetricList } from '../components/metrics';

const Metrics = () => {
    return (
        <div className="flex flex-col justify-center items-center py-5 px-10">
            <div className="w-2/4">
                <div>{/* bar chart */}</div>
                <div>
                    <MetricList />
                </div>
            </div>
        </div>
    );
};

export default Metrics;
