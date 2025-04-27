# OpenTibia Gateway

---

# Architecture Overview

```plaintext
+-------------------------+
|      External Clients   |
| (HTTP clients, browsers)|
+-----------+-------------+
            |
    HTTP Server (cmd/internal-api)
            |
+-----------v------------+
| internal/transport/http |
| - HTTP Handlers         |
| - HTTP Response Models  |
+-----------+------------+
            |
   gRPC Client to Core
            |
+-----------v------------+
| gRPC Server (cmd/core)   |
+-----------+------------+
            |
+-----------v------------+
| internal/transport/grpc |
| - gRPC Service Handlers |
| - gRPC Mappers          |
+-----------+------------+
            |
+-----------v------------+
|    internal/service     |
| - Business Logic        |
+-----------+------------+
            |
+-----------v------------+
|     internal/domain     |
| - Core Business Models  |
+-----------+------------+
            |
+-----------v------------+
| internal/provider/mysql |
| - Database Persistence  |
+-------------------------+
```

---

# Project Overview

**OpenTibia Gateway** is a service responsible for mediating communication between:
- A **gRPC backend server** (business logic and database access)
- An **HTTP API** (frontend-friendly API)

It acts as a gateway layer, decoupling database logic from external access and offering a clear, extensible and secure API surface.

**Main responsibilities:**
- Serve a gRPC API for internal trusted services.
- Serve a RESTful HTTP API for external clients (e.g., websites, bots, UIs).
- Maintain clear separation between transport, domain, and persistence layers.

---

# Key Design Choices

## 1. Layered Clean Architecture

| Layer | Purpose |
|:-----|:--------|
| cmd/ | Entry points (main.go for HTTP server, gRPC server) |
| internal/transport/http | HTTP routes, handlers, HTTP-specific models (DTOs) |
| internal/transport/grpc | gRPC handlers, protobuf mapping |
| internal/service | Business logic orchestration |
| internal/domain | Core domain models (e.g., Player) |
| internal/provider/mysql | Database layer (GORM entities, repositories) |

Each layer depends **only inward**, enforcing strict **dependency inversion** and isolation.

---

## 2. Why HTTP Models (DTOs) Are Separated

- **Security**: Only expose necessary fields to the outside world.
- **Versioning**: HTTP API can evolve independently from domain models.
- **Clarity**: HTTP response/request contracts are clear and explicit.
- **Decoupling**: Protect internal data shapes from leaking outside.

Example: `internal/transport/http/player_response.go` defines public-safe fields like ID, Name, Health, omitting internal fields like database IDs, conditions, etc.

---

## 3. Why gRPC + HTTP Split

- **gRPC**: Internal, high-performance service communication (trusted zone).
- **HTTP**: Public facing, user-friendly, frontend-consumable API (open zone).

This allows scaling horizontally, secure authentication internally, while keeping user APIs friendly.

---

## 4. Golang Idiomatic Patterns Used

| Concept | Application |
|:--------|:------------|
| Composition over Inheritance | Handlers compose services, not inherit |
| Dependency Injection | Handlers receive services and clients via constructors |
| Explicit mapping | Domain models are explicitly mapped to DTOs |
| Context propagation | `context.Context` is respected and passed |
| Error handling with status codes | gRPC errors mapped properly (NotFound, Internal, etc.) |
| Transport-layer separation | gRPC and HTTP have distinct mapping responsibilities |

---

# Folder Structure

```plaintext
cmd/
  core/           # gRPC Server (main.go)
  internal-api/   # HTTP Server (main.go)

internal/
  domain/         # Core business models (Player, etc.)
  provider/
    mysql/        # GORM database persistence
  service/        # Business logic orchestration
  transport/
    grpc/         # gRPC handlers and mapping
    http/         # HTTP handlers, HTTP DTOs (Response Models)

api/
  proto/v1/       # gRPC protobuf definitions (player.proto)

internal/protogen/v1/ # Generated gRPC Go files (player.pb.go, player_grpc.pb.go)

.vscode/
  launch.json     # Debug configurations for both gRPC and HTTP servers
```

---

# Running the project

### Prerequisites
- Go 1.21+
- Protoc compiler + protoc-gen-go + protoc-gen-go-grpc
- Make (optional but recommended)
- VSCode (with Go extensions recommended)

### 1. Compile Protobufs

```bash
make proto
```

This will generate gRPC Go files into `internal/protogen/v1/`.

### 2. Start gRPC Server

```bash
go run cmd/core/main.go
```

gRPC server listens on `localhost:50051`.

### 3. Start HTTP Server

```bash
go run cmd/internal-api/main.go
```

HTTP server listens on `localhost:8080`.

### 4. Example Call

```bash
curl http://localhost:8080/players/1
```

Response:
```json
{
  "id": 1,
  "name": "Knight",
  "level": 30,
  "vocation": 4,
  "health": 420,
  "health_max": 420,
  ...
}
```

---

# Debugging

### Debug Single Server
- Select `Debug Core (gRPC Server)` or `Debug Internal API (HTTP Server)` inside VSCode.

### Debug Both Servers Together
- Use compound configuration `Debug Both (Core + Internal API)` in VSCode Run and Debug panel.

.vscode/launch.json is already prepared for seamless experience.

---

# Deeper Dive: gRPC Error Handling

- Internal services return Go errors like `ErrNotFound`, `ErrInvalidArgument`.
- `internal/transport/grpc/error.go` maps those errors to proper gRPC status codes.
- HTTP server captures gRPC errors and maps them to correct HTTP responses:
  - gRPC `codes.NotFound` -> HTTP `404`
  - gRPC `codes.InvalidArgument` -> HTTP `400`
  - gRPC `codes.Internal` -> HTTP `500`

This standardizes API behavior and lets clients handle errors easily.

---

# Justification for Each Clean Architecture Principle

| Principle | Justification |
|:----------|:--------------|
| Dependency Inversion | HTTP and gRPC layers depend on service interfaces, not implementations |
| Single Responsibility | Transport layers only handle IO, services handle business logic |
| Open/Closed | Add new services or endpoints without modifying existing ones |
| Separation of Concerns | Domain models are isolated from transport and persistence |

✅ The project is aligned to production-level software architecture standards.

---

# Future Improvements (Roadmap)

- Add Authentication (gRPC and HTTP)
- Introduce environment-driven configuration (envconfig, viper)
- Introduce structured logging (Zap, Logrus)
- Implement retry/timeout strategies for gRPC clients
- Add graceful shutdown and context cancellation propagation
- Version the HTTP API (v1, v2, etc.)

---

# Conclusion

This project is designed to be **clean, professional, extensible**, and **scalable**.
It applies strong software engineering principles while remaining practical for real-world game server and API needs.

---

# Contacts

Built by Danilo Pucci.
Feel free to contribute, suggest improvements, or use this architecture as a solid base for your own production-grade Go services.

---

# License

[MIT License](LICENSE)

---

> "First make it work, then make it clean, then make it fast." — Kent Beck

