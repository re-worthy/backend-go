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
-   (not in runtime) https://github.com/pressly/goose (db migrations)

## Project Structure (TBD)
```
project-root/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
│
├── internal/                       # Private application code
│
├── pkg/                            # Public library code
│
├── .env.example
├── .gitignore
├── .local.db                       # local database file (sqlite3)
├── README.md
├── LICENSE
├── go.mod
└── go.sum
```
