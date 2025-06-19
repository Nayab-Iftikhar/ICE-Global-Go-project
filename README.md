# Todo App

This is a Todo application that allows users to create TodoItems, upload files, and store them in MySQL and S3 (via LocalStack). It uses Redis Streams to push notifications.

---

## Setup

1. **Clone the repo:**
   ```bash
   git clone <https://github.com/Nayab-Iftikhar/ICE-Global-Go-project.git>
   ```

2. **Start all services (MySQL/MariaDB, Redis, LocalStack, app) with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

---


## Makefile Commands

You can use the following `make` commands to simplify common development tasks:

- **`make run`**  
  Builds and starts the entire application stack (including the app, database, Redis, and LocalStack) using Docker Compose.

- **`make test`**  
  Runs all handler tests in the project.

- **`make benchmark`**  
  Executes all benchmarks in the `benchmark` directory, showing performance and memory allocation statistics.

- **`make migrate`**  
  Applies all database migrations to set up or update your local database schema.

These commands help automate setup, testing, benchmarking, and database migrations, making local development and CI/CD easier.

---




## Applying Migrations

1. **Install [golang-migrate](https://github.com/golang-migrate/migrate) if you don't have it:**
   ```bash
   brew install golang-migrate
   # or see https://github.com/golang-migrate/migrate#installation
   ```

2. **Apply migrations:**
   ```bash
   ./scripts/migrate.sh
   ```
   This will apply all migrations in the `migrations/` directory to your local database.

---

## Running Tests

- **Run all unit tests:**
  ```bash
  go test ./...
  ```

- **Run handler tests with verbose output:**
  ```bash
  go test -v ./handlers_test
  ```

---

## Running Benchmarks

- **Run all benchmarks:**
  ```bash
  go test -bench=. -benchmem -v ./benchmark
  ```

- **Or use the Makefile:**
  ```bash
  make benchmark
  ```

---






## Endpoints

- `POST /upload` — Upload a file
- `POST /todo` — Create a Todo item

See code and comments for more details.

---

## Notes

- Make sure Docker and Docker Compose are installed and running.
- LocalStack is used for S3 emulation; AWS credentials in `.env` should be test values.
- MariaDB is recommended for Apple Silicon (M1/M2/M3) compatibility.

---
