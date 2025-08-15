# Clean Architecture Structure

This project follows Clean Architecture principles, separating concerns into distinct layers with clear dependency directions.

## Architecture Layers

```
┌─────────────────────────────────────────────────────────────┐
│                    Frameworks & Drivers                    │
│                  (External Interfaces)                     │
├─────────────────────────────────────────────────────────────┤
│                   Interface Adapters                       │
│              (Controllers, Presenters, Gateways)           │
├─────────────────────────────────────────────────────────────┤
│                      Use Cases                             │
│                 (Application Business Rules)               │
├─────────────────────────────────────────────────────────────┤
│                      Entities                              │
│                 (Enterprise Business Rules)                │
└─────────────────────────────────────────────────────────────┘
```

## Directory Structure

```
internal/
├── domain/                     # Enterprise Business Rules
│   ├── entity/                 # Business entities
│   │   └── user.go            # User entity with business logic
│   └── repository/            # Repository interfaces
│       └── user_repository.go # User repository contract
├── usecase/                   # Application Business Rules
│   └── user_usecase.go       # User use cases/interactors
├── interface/                 # Interface Adapters
│   ├── controller/            # Controllers (handle HTTP)
│   │   └── user_controller.go # User HTTP handlers
│   └── presenter/             # Presenters (format output)
│       └── user_presenter.go  # User response formatting
└── infrastructure/            # Frameworks & Drivers
    ├── repository/            # Data access implementations
    │   └── memory_user_repository.go # In-memory user storage
    └── router/                # HTTP routing
        └── router.go         # Route configuration
```

## Layer Responsibilities

### 1. Domain Layer (`internal/domain/`)
- **Entities**: Core business objects with business rules
- **Repository Interfaces**: Contracts for data access
- **Independent**: No dependencies on external frameworks

### 2. Use Case Layer (`internal/usecase/`)
- **Application Business Rules**: Orchestrates business logic
- **Interactors**: Coordinates between entities and repositories
- **Depends on**: Domain layer only

### 3. Interface Adapters (`internal/interface/`)
- **Controllers**: Handle HTTP requests/responses
- **Presenters**: Format output data
- **Gateways**: Interface implementations
- **Depends on**: Use case and domain layers

### 4. Infrastructure Layer (`internal/infrastructure/`)
- **Repository Implementations**: Concrete data access
- **Web Framework**: HTTP server and routing
- **External Services**: Database, file system, etc.
- **Depends on**: All inner layers

## Dependency Flow

```
main.go → Infrastructure → Interface → UseCase → Domain
```

- Dependencies point inward (Dependency Inversion Principle)
- Inner layers don't know about outer layers
- Interfaces define contracts between layers

## Benefits

1. **Testability**: Easy to unit test business logic
2. **Maintainability**: Clear separation of concerns
3. **Flexibility**: Easy to change external dependencies
4. **Independence**: Business logic independent of frameworks
5. **Scalability**: Easy to add new features following established patterns

## Key Principles Applied

- **Single Responsibility**: Each layer has one reason to change
- **Open/Closed**: Open for extension, closed for modification
- **Dependency Inversion**: Depend on abstractions, not concretions
- **Interface Segregation**: Small, focused interfaces
- **Don't Repeat Yourself**: Shared logic in appropriate layers

## Running the Application

```bash
# Build the application
go build -o app .

# Run the application
./app

# Or run directly
go run main.go
```

## API Endpoints

- `GET /health` - Health check
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/{id}` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/{id}` - Update user
- `DELETE /api/v1/users/{id}` - Delete user

## Adding New Features

1. **Entity**: Add new business entity in `domain/entity/`
2. **Repository**: Define repository interface in `domain/repository/`
3. **Use Case**: Implement business logic in `usecase/`
4. **Controller**: Add HTTP handlers in `interface/controller/`
5. **Presenter**: Format responses in `interface/presenter/`
6. **Infrastructure**: Implement repository in `infrastructure/repository/`
7. **Router**: Register routes in `infrastructure/router/`
8. **Main**: Wire dependencies in `main.go`
