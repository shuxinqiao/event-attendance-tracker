import React, { useState } from 'react';
import Events from './Events';
import Login from './Login';
import 'primereact/resources/themes/saga-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';

function App() {
	const [user, setUser] = useState(null);

	return (
		<div className="App">
			<h1>Event Attendance System</h1>
			{user ? <Events userId={user.id} /> : <Login setUser={setUser} />}
		</div>
	);
}

export default App;