args := $(wordlist 2, 100, $(MAKECMDGOALS))
tempFile = ".temp"

build:
	@touch ${tempFile}
	@/usr/bin/time -o ${tempFile} go build -v -o bin/bin cmd/api/main.go
	@echo "Build complete: ./bin/bin\nTime spent:"
	@cat ${tempFile}
	@rm ${tempFile}

dev:
	@ENVIRONMENT=DEVELOPMENT go run cmd/api/main.go

env_init:
	@if [ -e ".env" ] || [ -e ".env.dev" ]; then \
		read -p "Warning! \".env\" or \".env.dev\" file found. Override them? [y/N]: " ans; \
		if [ ! "$$ans" = "y" ]; then \
			exit 1; \
		fi \
	fi
	@echo
	@echo "# Production env (overrided by .dev)" > .env
	@echo "# Development env" > .env.dev
	@echo "\n# Copied from .env.example dont forget to change values\n\n" >> .env
	@echo "\n# Copied from .env.example dont forget to change values\n\n" >> .env.dev
	@cat ./.env.example >> .env
	@cat ./.env.example >> .env.dev
	@echo "ENVIRONMENT=\"PRODUCTION\"" >> .env
	@echo "ENVIRONMENT=\"DEVELOPMENT\"" >> .env.dev
	@echo ".env, .env.dev are created.\nContent from .env.example is copied into them"

test:
	@go test ./pkg/... ./internal/... ./cmd/... $(FLAGS)

init:
	@make env_init;

sqlc:
	@sqlc generate -f internal/db/sqlc/sqlc.yaml

goose:
	@GOOSE_DRIVER=turso GOOSE_DBSTRING=file:./.local.db goose -dir "./internal/db/migrations/" $(args) 

default:
	dev;

%::
	@echo ""
