# Split The Bill - Server

[![Go Webserver Testing](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml/badge.svg)](https://github.com/lab-64/split-the-bill-server/actions/workflows/go.yml)

---
# Endpoints

**Register User** <br />
-> POST [$URL/api/user/register](http://localhost:8080/api/user/register) <br />
Body: { <br />
&nbsp;&nbsp;&nbsp; "email": "", <br />
&nbsp;&nbsp;&nbsp; "password": "", <br />
&nbsp;&nbsp;&nbsp; "confirmationPassword": "" <br />
} <br />
Response: { <br />
&nbsp;&nbsp;&nbsp; "message": "", <br />
&nbsp;&nbsp;&nbsp; "status": "error/ok", <br />
&nbsp;&nbsp;&nbsp; "user": "user email addr", <br />
}


**Login User** <br />
-> POST [$URL/api/user/login](http://localhost:8080/api/user/login) <br />
Body: { <br />
&nbsp;&nbsp;&nbsp; "email": "", <br />
&nbsp;&nbsp;&nbsp; "password": "", <br />
} <br />
Response: { <br />
&nbsp;&nbsp;&nbsp; "cookieAuth": { <br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; "id": "", <br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; "username": "", <br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; "expirationTime": "", <br />
&nbsp;&nbsp;&nbsp; }, <br />
&nbsp;&nbsp;&nbsp; "status": "error/ok", <br />
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
