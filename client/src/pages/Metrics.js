import React from 'react';
import { MetricList } from '../components/metrics';
import Chart from 'react-google-charts';

const Metrics = () => {
    return (
        <div className="flex flex-col justify-center items-center py-5 px-10">
            <div className="w-2/4">
                <div>
                    <Chart
                        className="w-full"
                        height={'300px'}
                        chartType="Bar"
                        loader={<div>Loading Chart</div>}
                        data={[
                            ['', 'Task Completed'],
                            ['2014', 1000],
                        ]}
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
