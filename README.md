## Hexagonal Architecture GO

```
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
│   │   ├── product_service.go
│   │
│   ├── adapters/            # Adapters (ports)
│   │   ├── db/
│   │   │   └── gorm_product_repository.go
│   │   ├── http/
│   │   │   └── product_handler.go
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

```



## Author

- [floxydio](https://www.github.com/floxydio)

