# Go API

This application is a project to learn about [GIN](https://gin-gonic.com/) framework. It also followed design patterns like:
- Domain-driven design (DDD).
- Dependency Injection (DI).
It also followed SOLID principle.
It has coverted with unit tests [testify](https://github.com/stretchr/testify) and [sqlmock](https://pkg.go.dev/github.com/data-dog/go-sqlmock)
[JWT](https://github.com/golang-jwt/jwt) has been implemented

## Installation

To install and run the Go API locally, follow these steps:

1. Clone the repository:

```shell
git clone https://github.com/yakob-abada/go-api.git
```

2. Change into the project directory:

```shell
cd go-api
```

3. Build the project and run

```shell
docker compose up
```

The API should now be running on `http://localhost:8080`.

## Gym application
- It exposes come Restful APIs that ables loggedin gym members to list all active sessions for given week and join them if it's not full.
    - GET `http://localhost:8080/active-session` -> to get active session for given week.
    - POST `http://localhost:8080/login` -> to login and retrieve JWT token.
        - body 
        ```
            {
                "username": "yakob.abada",
                "password": "secret"
            }
        ```
    - POST `http://localhost:8080/session/{session_id}/join`
        - header ```Authorization: Breaer {JWT token} ```
## Usage

[Explain how to use the API, including any necessary authentication or request formats.]

## Run tests

```shell
make tests
```

## Things to improve.
- Adding error logs for production purposes.
- Improve mysql message to make it more user friendly.
- Handle query injection.
- Add a cron job that creates session at the beginning of every week.
- User signup.
- Refresh token