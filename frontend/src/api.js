import axios from 'axios';

// Create an axios instance with default settings
const api = axios.create({
    baseURL: 'http://192.168.2.160:80/api/',
    // baseURL: 'http://localhost:80/api/',
    header: {
        'Content-Type': 'application/json',
    },
    withCredentials: true,
});

// Interceptor for logging or handling errors globally
api.interceptors.response.use(
    (response) => response,
    (error) => {
        // Log or handle errors gloablly
        console.error("API error:", error);
        return Promise.reject(error);
    }
);

// Fetch events
export const fetchEvents = async () => {
    try {
        const response = await api.get('/events');
        return response.data;
    } catch (error) {
        console.error("Failed to fetch events:", error);
        throw error;
    }
};

// Send POST to check in a user for events
export const checkInUser = async (userId, eventId) => {
    try {
        const response = await api.post('/checkin', { userId, eventId });
        return response.data;
    } catch (error) {
        console.error("Failed to check in user:", error);
        throw error;
    }
};

export default api;