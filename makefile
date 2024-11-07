# Makefile

# Whole Project Level
# Start whole project docker
start-project:
	sudo docker-compose up -d

# Stop whole project docker
stop-project:
	sudo docker-compose down

# Restart whole project docker
restart-project:
	sudo docker-compose down
	sudo docker-compose up --build -d

# Reset whole project docker
reset-project:
	sudo docker-compose down -v
	sudo docker-compose up --build -d


# PostgreSQL
# Start PostgreSQL container
start-db:
	sudo docker-compose up -d postgres

# Stop PostgreSQL container
stop-db:
	sudo docker-compose down postgres

# Rebuild PostgreSQL with fresh data
reset-db:
	sudo docker-compose down -v postgres
	sudo docker-compose up -d postgres

# Run psql command to access the database shell
db-shell:
	sudo docker exec -it event_attendance_db psql -U postgres -d event_attendance
