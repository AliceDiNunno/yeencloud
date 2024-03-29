import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {CurrentSetupStep} from "../../api/setup";

export default (): JSX.Element => {
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await CurrentSetupStep();
                navigate(`/setup/${response}`);
            } catch (err) {
                console.error(err);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, [navigate]);

    return (
        <div>
            {loading ? 'Loading...' : 'Finished loading'}
        </div>
    );
};
