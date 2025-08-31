## Hexagonal Architecture GO

`
/app
│── cmd/
│   └── app/
│       └── main.go          # Entry point (setup Fiber, DI, dll.)
│
│── internal/
│   ├── domain/              # Core business logic (Entities + Ports)
│   │   ├── user.go          # Entity: User
│   │   └── ports.go         # Ports (interface service/repo)
│   │
│   ├── application/         # Use cases (Application services)
│   │   ├── booking_service.go
│   │   └── user_service.go
│   │
│   ├── adapters/            # Adapters (ports)
│   │   ├── db/
│   │   │   ├── gorm_booking_repository.go
│   │   │   └── gorm_user_repository.go
│   │   ├── http/
│   │   │   ├── booking_handler.go
│   │   │   └── user_handler.go
│   │   └── config/
│   │       └── db.go        # DB
│   │
│   └── utils/
│       └── response.go      # Custom Response
│
│── pkg/                     # Library
│   └── logger/
│       └── logger.go
│
│── go.mod
│── go.sum

`