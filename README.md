# Split The Bill - Server

[![Go Webserver Testing](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml/badge.svg)](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml)

---
# Start Application
We can start the application with different storage types.

## Ephemeral Storage

Add the following code in the "select storage" section in `main.go`:
```go
// Select storage
storage := ephemeral.NewEphemeral()
```
Start the application with:
```shell
make run
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
# Endpoints

**Register User** <br />
-> POST [$URL/api/user/register](http://localhost:8080/api/user/register) <br />
Body: { <br />
&nbsp;&nbsp;&nbsp; "username": "", <br />
&nbsp;&nbsp;&nbsp; "password": "", <br />
} <br />
Response: { <br />
&nbsp;&nbsp;&nbsp; "message": "", <br />
&nbsp;&nbsp;&nbsp; "status": "error/ok", <br />
&nbsp;&nbsp;&nbsp; "User": "username", <br />
}


**Login User** <br />
-> POST [$URL/api/user/login](http://localhost:8080/api/user/login) <br />
Body: { <br />
&nbsp;&nbsp;&nbsp; "username": "", <br />
&nbsp;&nbsp;&nbsp; "password": "", <br />
} <br />
Response: { <br />
&nbsp;&nbsp;&nbsp; "cookieAuth": { <br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; "UserID": "", <br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; "Token": "", <br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; "ValidBefore": "", <br />
&nbsp;&nbsp;&nbsp; }, <br />
&nbsp;&nbsp;&nbsp; "status": "error/ok", <br />
} <br />