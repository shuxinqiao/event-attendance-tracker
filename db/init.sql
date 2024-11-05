-- Add users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100), -- Password is required only for admins and super admins
    role VARCHAR(20) NOT NULL CHECK (role IN ('super_admin', 'admin', 'user')), -- Enforce specific roles
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add events table
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    date DATE NOT NULL,
    location VARCHAR(100) NOT NULL
);

-- Add checkins table with references to users and events
CREATE TABLE IF NOT EXISTS checkins (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    event_id INT REFERENCES events(id),
    checkin_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert super_admin with a default password
INSERT INTO users (username, password, role) VALUES
    ($SUPERADMIN_USERNAME, $SUPERADMIN_PASSWORD, 'super_admin')
ON CONFLICT DO NOTHING;