import React, { useEffect, useState } from 'react';
import CheckIn from './CheckIn';
import "./Event.css"
import { fetchEvents } from './api';

import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Sidebar } from 'primereact/sidebar';
import { Button } from 'primereact/button';


const Events = ({ userId }) => {
    const [events, setEvents] = useState([]);
    const [visible, setVisible] = useState(false);

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
        <div className='event-page-wrap'>
            <div className='head-wrap'>
                <h2 className='head-title'>Events</h2>
                <Sidebar visible={visible} position="left" onHide={() => setVisible(false)}>
                    <h2>Sidebar</h2>
                    <p>
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
                        Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
                    </p>
                </Sidebar>
                <Button className='head-nav-btn' icon="pi pi-arrow-right" onClick={() => setVisible(true)} />
            </div>
            
            <DataTable value={events}>
                <Column field='name' header="Event Name" />
                <Column field='date' header="Date" />
                <Column field='location' header="Location" />
                <Column header="Check In" body={(rowData) => <CheckIn eventId={rowData.id} userId={userId} />} />
            </DataTable>
            
            
        </div>
    );
};

export default Events