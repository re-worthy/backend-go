## Git hooks

```bash
# sets git-hooks local path
git config --local core.hooksPath .git-hooks

# makes files executable
chmod +x ./.git-hooks/pre-commit
chmod +x ./.git-hooks/commit-msg
```

### Prefixes:
-   [HF] - hot fix
-   [F] - Feature
-   [BF] - bug fix
-   [DX] - improved Developer eXperience
-   [R] - refactor
-   [CUSTOM] - something unknown but so much wanted

## Stack
-   go-playground/validator (json validation)
-   github.com/caarlos0/env (env variables validation)
-   https://sqlc.dev (sql compiler (sql typechecker -> golang funcs))
-   [Future] golang.org/x/crypto/argon2 (password hashing)
-   [Future] https://github.com/pressly/goose (db migrations)

## Project Structure (potential)
```
project-root/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
│
├── internal/                       # Private application code
│   ├── config/
│   │   └── config.go              # Configuration structures and loading
│   │
│   ├── server/
│   │   ├── server.go              # Server setup and initialization
│   │   ├── routes.go              # Route definitions
│   │   ├── middleware.go          # Custom middleware
│   │   └── handlers/              # HTTP handlers
│   │       ├── health.go
│   │       ├── user.go
│   │       └── common.go
│   │
│   ├── models/                    # Data models
│   │   ├── user.go
│   │   └── response.go
│   │
│   ├── repository/                # Database interactions
│   │   ├── interfaces.go
│   │   ├── user_repository.go
│   │   └── postgres/
│   │       └── postgres.go
│   │
│   └── service/                   # Business logic
│       ├── interfaces.go
│       └── user_service.go
│
├── pkg/                           # Public library code
│   ├── logger/
│   │   └── logger.go
│   │
│   └── utils/
│       ├── http.go
│       └── validator.go
│
├── migrations/                    # Database migrations
│   ├── 001_initial.up.sql
│   └── 001_initial.down.sql
│
├── scripts/                       # Build and deployment scripts
│   ├── build.sh
│   └── deploy.sh
│
├── configs/                       # Configuration files
│   ├── config.yaml
│   └── config.production.yaml
│
├── tests/                        # Integration and e2e tests
│   ├── integration/
│   │   └── user_test.go
│   └── mocks/
│       └── user_service_mock.go
│
├── api/                          # API documentation
│   └── swagger.yaml
│
├── Dockerfile                    # Docker configuration
├── docker-compose.yml            # Docker compose configuration
├── .gitignore
├── README.md
├── go.mod
└── go.sum
```
