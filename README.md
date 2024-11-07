# event-attendance-tracker
A web-based event attendance system with user login, admin controls, and check-in functionality, built with Go, React, and Docker.

## Contents
- [Setup Procedure](#setup-procedure) - Setup for your specific settings
- [Scripts](#script) - Guide for a easy and quick management


## Setup Procedure

### Database Setup

This project uses Docker and Docker-compose to set up a PostgreSQL database. Follow these instructions to configure the database with environment-specific credentials.

#### Step 1: Create a `.env` File

In the project root, create a `.env` file with the following variables:

```env
export POSTGRES_USER=your_postgres_user
export POSTGRES_PASSWORD=your_postgres_password
export POSTGRES_DB=your_database_name
export SUPERADMIN_USERNAME=superadmin
export SUPERADMIN_PASSWORD=superadminpassword
```

- Quick Tip: Use `"'var_name'"` to have `'var_name'`in real script.

#### Step 2: Substitude Placeholder by Your `.env` File Setting

```bash
POSTGRES_USER --> username of db admin (complete permission)
POSTGRES_PASSWORD --> password of db admin
POSTGRES_DB --> name of db
SUPERADMIN_USERNAME --> name of super admin user (for account management)
SUPERADMIN_PASSWORD --> password of super admin user
```

#### Step 3: Start Postgres Docker

```bash
$ make start-db
```

More details about easier control over postgres container see [script](#script).


## Script

### Using the Makefile for Docker Management

This project’s Makefile provides convenient commands for managing Docker containers for the entire project and the PostgreSQL database separately. Below are the commands and how to use them.

### Whole Project Commands

These commands control all containers in the Docker Compose setup, including Nginx, the Go backend, and PostgreSQL.

1. **Start the Project**:
   - Starts all containers in the project.
   ```bash
   $ make start-project
   ```

2. **Stop the Project**:
   - Stops all containers in the project.
   ```bash
   $ make start-project
   ```

3. **Restart the Project**:
   - Stops and removes all containers ***without touching mount data***, rebuilds them, and starts them up with the latest changes.
    ```bash
    $ make restart-project
    ```

3. **Reset the Project**:
   - Stops and removes all containers ***and their mount data***, rebuilds them, and starts them up with the latest changes.
    ```bash
    $ make reset-project
    ```

### PostgreSQL-Specific Commands

These commands only control the PostgreSQL container, useful if you need to restart or reset the database independently of the other containers.

1. **Start PostgreSQL**:
    - Starts only the PostgreSQL container.
    ```bash
    $ make start-db
    ```

2. **Stop PostgreSQL**:
    - Stops only the PostgreSQL container.
    ```bash
    $ make stop-db
    ```

3. **Rebuild PostgreSQL with Fresh Data**:
    - Stops the PostgreSQL container, removes any existing data, and starts a fresh instance with the initial data.
    ```bash
    $ make reset-db
    ```

4. **Access the PostgreSQL Shell**:
    - Opens an interactive psql shell for accessing the PostgreSQL database directly. Use this to run SQL commands manually.
    ```bash
    $ make db-shell
    ```

### Important Notes

- **Permissions**: All commands use `sudo` to ensure Docker commands have the necessary permissions. You may modify `sudo` if your Docker setup does not require it.
- **Rebuilding Containers**: Use the `reset-project` and `reset-db` commands carefully, as they will recreate the containers and potentially reset the database.
- **Environment Configuration**: Ensure the `.env` file is configured correctly with your environment variables before starting the project.
