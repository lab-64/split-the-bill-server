# Split The Bill - Server

[![Go Webserver Testing](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml/badge.svg)](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml)

---
# Developer Get Started
**1. Install Go**
- follow the instructions: https://go.dev/doc/install

**2. Clone the repository**

**3. Install [Reflex](https://github.com/cespare/reflex) package (needed for Hot Reload)**
- run `go install github.com/cespare/reflex@latest` 

**4. Install [swag](https://github.com/swaggo/swag) package (needed for fiber-swagger)**
- run `go install github.com/swaggo/swag/cmd/swag@latest`
---

# Start Application
**1. Set storage type in `.env`:**
- `STORAGE_TYPE=ephemeral` for Ephemeral
- `STORAGE_TYPE=postgres` for Postgres

**2a. Start the application (docker, ephemeral & postgres) with:**
```shell
make start-postgres
```

**2b. (OR) Start the application (no docker, ephemeral only) with:**
```shell
make watch
```
**3. Stop the application with:**
```shell
make stop-postgres
```

**4. Reset the database with:**
```shell
make reset-db
```

---
# Testing

```shell
make test-all
```

---

# URLs

- Swagger API: http://localhost:8080/swagger/ 

