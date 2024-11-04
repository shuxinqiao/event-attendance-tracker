import React, { useState } from "react";
import api from "./api";

const CheckIn = ({ eventId }) => {
    const [userId, setUserId] = useState('');

    const handleCheckIn = async () => {
        try {
            const response = await api.post('/checkin', { userId, eventId });
            alert('Check-in successful!');
        } catch (error) {
            console.error('Check-in failed:', error);
            alert('Check-in failed');
        }
    };

    return (
        <div>
            <h2>Check In</h2>
            <input type="text" placeholder="User ID" value={userId} onChange={(e) => setUserId(e.target.value)} />
            <button onClick={handleCheckIn}>Check In</button>
        </div>
    );
};

export default CheckIn;