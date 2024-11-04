import React, { useState } from "react";
import api from "./api";

const CheckIn = ({ userId, eventId }) => {
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
            <button onClick={handleCheckIn}>Check In</button>
        </div>
    );
};

export default CheckIn;