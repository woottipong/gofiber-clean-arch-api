# Go HTTP API Clean Architecture
This is a sample RESTful API built with GoFiber and following the Clean Architecture pattern.

## Directory Structure
```go
.
├── cmd
│   └── api
│       └── main.go
├── config
│       └── config.go
├── internal
│   ├── handler
│   │   ├── user.go
│   ├── entity
│   │   └── user.go
│   ├── repository
│   │   └── user.go
│   ├── usecase
│   │   └── user.go
│   └── infrastructure
│       ├── database.go
│       └── fiber.go
└── go.mod
```

- `cmd/api/main.go`: entry point of the application
- `config`: holds the configuration code for the application.
- `handler`: holds the code for handling HTTP requests, responses and setup router.
- `entity`: holds the code for the application's entities, which represent the core business logic of the application.
- `repository`: holds the code for accessing the database and managing data persistence.
- `usecase`: holds the code for implementing the business logic of the application.
- `infrastructure`: holds the code for initializing the database and web framework.

## Usage
1. Clone this repository
2. `cd gofiber-clean-arch-api`
3. `docker-compose up -d`
4. Send HTTP requests to `[POST] localhost:5000/v1/users`

## API Endpoints
- `GET /users`: get all users
- `POST /users`: create a new user
- `GET /users/:id`: get a user by ID
- `PUT /users/:id`: update an existing user by ID
- `DELETE /users/:id`: delete a user by ID

## Dependencies
- `Fiber`: web framework
- `go-sql-driver/mysql`: MySQL driver for Go
- `validator`: github.com/go-playground/validator/v10