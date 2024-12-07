tempFile = ".temp"


build:
	@touch ${tempFile}
	@/usr/bin/time -o ${tempFile} go build -v -o bin/bin cmd/api/main.go
	@echo "Build complete: ./bin/bin\nTime spent:"
	@cat ${tempFile}
	@rm ${tempFile}

dev:
	@go run cmd/api/main.go


test:
	@go test ./pkg/... ./internal/... ./cmd/... $(FLAGS)
