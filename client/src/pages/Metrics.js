import { React /*, useEffect, useState*/ } from 'react';
import { MetricList } from '../components/metrics';
//import axios from 'axios';
import Chart from 'react-google-charts';

var arr = [['', 'Task Completed']];
var days = 7;

for (let i = days; i > 0; i--) {
    let d = new Date(new Date().getTime() - 24 * i * 3600 * 1000).toLocaleDateString();
    arr.push([d, 1000]);
}

console.log(arr);

const Metrics = () => {
    return (
        <div className="flex flex-col justify-center items-center py-5 px-10">
            <div className="w-3/5">
                <div>
                    <Chart
                        className="w-full"
                        height={'300px'}
                        chartType="Bar"
                        loader={<div>Loading Chart</div>}
                        data={arr}
                        options={{
                            legend: { position: 'none' },
                            chart: {
                                title: 'Past Activity',
                                subtitle: 'Stuff from the last week',
                            },
                        }}
                    />
                </div>
                <div>
                    <MetricList />
                </div>
            </div>
        </div>
    );
};

export default Metrics;
