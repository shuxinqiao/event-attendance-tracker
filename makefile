# Makefile

# Start PostgreSQL container
start-db:
	sudo docker-compose up -d postgres

# Stop PostgreSQL container
stop-db:
	sudo docker-compose down

# Rebuild PostgreSQL with fresh data
reset-db:
	sudo docker-compose down -v
	sudo docker-compose up -d postgres

# Run psql command to access the database shell
db-shell:
	sudo docker exec -it event_attendance_db psql -U postgres -d event_attendance
