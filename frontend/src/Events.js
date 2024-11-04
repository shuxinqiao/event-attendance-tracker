import React, { useEffect, useState } from 'react';
import { fetchEvents } from './api';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import CheckIn from './CheckIn';

const Events = () => {
    const [events, setEvents] = useState([]);

    useEffect(() => {
        const loadEvents = async () => {
            try {
                const data = await fetchEvents();
                setEvents(Array.isArray(data) ? data : []);
            } catch (error) {
                console.error("Error loading events:", Error);
                setEvents([]);
            }
        };

        loadEvents();
    }, []);

    return (
        <div>
            <h2>Events</h2>
            <DataTable value={events}>
                <Column field='name' header="Event Name" />
                <Column field='date' header="Date" />
                <Column field='location' header="Location" />
                <Column header='Check In' body={(rowData) => <CheckIn eventId={rowData.id} />} />
            </DataTable>
        </div>
    );
};

export default Events