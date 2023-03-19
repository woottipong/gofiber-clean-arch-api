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
│   ├── router
│   │   └── user.go
│   └── infrastructure
│       ├── database.go
│       └── fiber.go
└── go.mod
```
- `cmd/api/main.go`: entry point of the application
- `config`: holds the configuration code for the application.
- `handling`: holds the code for handling HTTP requests and responses.
- `entity`: holds the code for the application's entities, which represent the core business logic of the application.
- `repository`: holds the code for accessing the database and managing data persistence.
- `usecase`: holds the code for implementing the business logic of the application.
- `infrastructure`: holds the code for initializing the database and web framework.
The cmd directory still holds the main.go file, which is responsible for setting up the dependencies an