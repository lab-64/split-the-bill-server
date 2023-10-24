# Split The Bill - Server

[![Go Webserver Testing](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml/badge.svg)](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml)

---
# Developer Get Started
**1. Install Go**
- follow the instructions: https://go.dev/doc/install

**2. Clone the repository**

**3. Install [Reflex](https://github.com/cespare/reflex) package (needed for Hot Reload)**
- run `go install github.com/cespare/reflex@latest` in the terminal 

**4. Install [swag](https://github.com/swaggo/swag) package (needed for fiber-swagger)**
- run `go install github.com/swaggo/swag/cmd/swag@latest`


---

# Start Application

## Ephemeral Storage

Add the following code in the "select storage" section in `main.go`:
```go
// Select storage
storage := ephemeral.NewEphemeral()
```
Start the application with:
```shell
make watch
```

## Postgres Database

Add the following code in the "select storage" section in `main.go`:
```go
// Select storage
storage, err := database.NewDatabase()
if err != nil {
    log.Fatal(err)
}
```

Start the application with:
```shell
make start-postgres
```
End the application with:
```shell
make stop-postgres
```
Reset the database with:
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

