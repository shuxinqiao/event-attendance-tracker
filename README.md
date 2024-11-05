# event-attendance-tracker
A web-based event attendance system with user login, admin controls, and check-in functionality, built with Go, React, and Docker.


### Database Setup

This project uses Docker and Docker-compose to set up a PostgreSQL database. Follow these instructions to configure the database with environment-specific credentials.

#### Step 1: Create a `.env` File

In the project root, create a `.env` file with the following variables:

```env
export POSTGRES_USER=your_postgres_user
export POSTGRES_PASSWORD=your_postgres_password
export POSTGRES_DB=your_database_name
export SUPERADMIN_USERNAME="'superadmin'"
export SUPERADMIN_PASSWORD="'superadminpassword'"
```

Use "'var_name'" to have 'var_name' in real script.

#### Step 2: Substitude Placeholder by Your `.env` File Setting

```sh
$ source .env
$ envsubst < db/init.sql > db/init_actual.sql
```

#### Step 3: Start Postgres Docker

```sh
$ make start-db
```

More details about easier control over postgres container see [db script](#db-script).