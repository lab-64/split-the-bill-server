# Split The Bill - Server

[![Go Webserver Testing](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml/badge.svg)](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml)

---
# Endpoints

Register User <br />
-> [$URL/api/register](http://localhost:8080/api/user/register) <br />
Body: { <br />
    "email": "", <br />
    "password": "", <br />
    "confirmationPassword": "" <br />
} <br />
Response: {
    "message": "",
    "status": "error/ok",
    "user": "user email addr",
}

---
# Start Application

```shell
docker-compose up -d --build
```

The postgresql database and the webserver will be started.

The webserver is accessable under: http://localhost:8080/

# End Application

```shell
docker-compose down
```

# Reset Postgres Database

```shell
docker-compose down --volumes
```