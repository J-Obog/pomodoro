import { React, useState, useEffect, useContext } from 'react';
import { AuthContext } from '../../context/AuthContext';
import Metric from './Metric';
import axios from 'axios';

const MetricList = () => {
    const { access } = useContext(AuthContext);
    const [metrics, setMetrics] = useState([]);

    useEffect(() => {
        axios
            .get(`${process.env.REACT_APP_API_URL}/user/metrics`, { headers: { Authorization: access } })
            .then(({ data }) => {
                setMetrics(data.metrics);
            })
            .catch((err) => {});
    }, [access, setMetrics]);

    return (
        <div className="bg-gray-100 rounded-lg p-5">
            <div>
                {metrics.map((m) => (
                    <Metric key={m.metric.replace(' ', '-')} {...m} />
                ))}
            </div>
        </div>
    );
};

export default MetricList;
