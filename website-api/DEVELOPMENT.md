# Development guide for website-api

This document records conventions and examples for developing the `website-api` service.
Keep it concise and practical so developers can onboard quickly.

## Goals
- Keep handlers thin: parsing, validation, and HTTP concerns only.
- Keep services stateless: business logic only, constructed once at app start.
- Keep repositories framework-agnostic: accept `context.Context` and read the request-scoped DB from it.
- Avoid Gin dependency in services/repositories; use `context.Context` instead.

## Project layout (recommended)

```
website-api/
├─ main.go
├─ go.mod
├─ .env.example
├─ database/            # DB init & migrations helpers
├─ models/              # GORM models (DB schema)
├─ repositories/        # DB access implementations per module
│  └─ users/            # e.g. repositories/users
├─ services/            # business logic per module
│  └─ users/
├─ handlers/            # HTTP handlers per module
│  └─ users/
├─ middlewares/         # Gin middlewares (DB context, auth, etc.)
├─ internal/            # internal shared types (errors, pagination)
└─ docs/
```

## Types and mapping conventions
- DB models: `models/*.go` (GORM models). Do not use these structs as transport DTOs.
- Service/domain types: `services/<module>/types.go` — types used inside services and returned to handlers.
- Transport DTOs: `handlers/<module>/dto.go` — request/response types with Gin binding tags.
- Mapping: provide small mapper functions in `services/<module>/mapper.go` or `handlers/<module>/mappers.go` to convert between models, domain types, and DTOs. Keep mapping explicit.

## Middleware: DB context
- Register a DB middleware that:
  - Attaches a *gorm.DB bound to the request context in the Gin `Context` (for handlers).
  - Also stores the *gorm.DB in the `request.Context` under a typed key so repositories/services can access it via `context.Context`.
- Use typed context keys to avoid collisions.

Example (existing): `middlewares/db_middleware.go` stores `db` in both `gin.Context` and `request.Context`.

Preferred helper functions (already present):

- `GetDBFromContext(c *gin.Context) *gorm.DB` — convenience for handlers.
- `DBFromContext(ctx context.Context) *gorm.DB` — repositories/services call this to obtain the request DB.

Important: `DBFromContext(ctx)` should return a *gorm.DB that already has the provided `ctx` attached, so repositories do not need to call `WithContext(ctx)` again.

## Repository pattern
- Place repository interfaces and implementations under `repositories/<module>`.
- Repository methods accept `context.Context` and read the DB with `middlewares.DBFromContext(ctx)`.
- Repositories should not import Gin.
- Return domain models (or `models.*`) and let services map them to domain types.

Example interface:

```go
type Repository interface {
    GetByID(ctx context.Context, id uint) (*models.User, error)
    Create(ctx context.Context, u *models.User) error
}
```

Implementation note: assume `DBFromContext(ctx)` returns a DB with `ctx` attached and call `db.First(&m, id)` directly.

## Service pattern
- Services are constructed once in `main.go` with their repository dependencies.
- Service methods accept `context.Context` and call repository methods which will obtain the request-scoped DB from `ctx`.
- Keep services free of Gin and *gorm.DB fields.

Example service signature:

```go
type Service interface {
    GetCurrent(ctx context.Context, id uint) (*User, error)
}
```

## Handlers
- Handlers are mounted in `handlers/<module>/handlers.go` via a `RegisterRoutes(rg *gin.RouterGroup, svc Service)` function.
- Handlers extract the request DB only when needed using `middlewares.GetDBFromContext(c)` or pass `c.Request.Context()` to services (services use DBFromContext internally).
- Prefer letting services read the DB from `context.Context` — handler simply forwards `c.Request.Context()`.

Example registration in `main.go`:

```go
userRepo := repositories.users.New()
userSvc := services.users.New(userRepo)
handlers.users.RegisterRoutes(v1.Group("/users"), userSvc)
```

Handler example (passing context only):

```go
func getMe(c *gin.Context, svc users.Service) {
    id := /* extract from auth */
    u, err := svc.GetCurrent(c.Request.Context(), id)
    // handle response
}
```

Alternative: handler can call `middlewares.GetDBFromContext(c)` and pass the DB explicitly to a service method that accepts `*gorm.DB` as parameter.

## Transactions
- Start transactions in handlers or services depending on the scope:
  - If multiple services/repos must share a tx, start it in the handler and put the tx DB into the request context so downstream calls pick it up.
  - Provide helpers: `BeginTx(ctx) (context.Context, *gorm.DB, func(err error) error)` to begin/commit/rollback and return a context containing the tx.

## Testing
- Unit test repositories using an in-memory DB or a test database; inject a context that contains a *gorm.DB.
- Unit test services by mocking repositories (define small mock implementations).
- Test handlers using Gin's httptest utilities and provide a middleware that injects a test DB into the request context.

## Examples / patterns to follow
- Services constructed once in `main.go`.
- Repositories read DB from `context.Context`.
- Handlers are thin: bind DTOs, call service methods, return responses.

## Quick checklist for new module
- [ ] models/<module> (if DB-backed)
- [ ] repositories/<module>/repo.go + repo_test.go
- [ ] services/<module>/service.go + service_test.go
- [ ] handlers/<module>/handlers.go + dto.go + handlers_test.go
- [ ] register routes in `main.go`

## Reference: useful snippets
- `middlewares/db_middleware.go` — ensures `gin.Context` and `request.Context` contain the request-scoped DB.
- Use `context.Context` as the single carrier of request-scoped values in services/repositories.

## Notes
- Keep exported interfaces small and focused to make unit-testing easy.
- Prefer explicitness: pass `context.Context` and let lower layers obtain the DB. Avoid passing `*gin.Context` into services.

---

If you want, I can scaffold one module (e.g., `users`) following this guide and wire it into `main.go` as a demonstration. Please tell me which module to scaffold.
