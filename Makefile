LDFLAGS_STATIC=-ldflags '-extldflags "-static"'

.PHONY: format

format:
	@go fmt ./...

build:
	@echo "Building static app executable..."
	@mkdir -p ./bin
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a ${LDFLAGS_STATIC} -o bin/main ./cmd/main

gen-mock:
	mockgen -source=internal/redis/redis.go -destination=internal/tools/testing/redis/mock_redis.go
	mockgen -source=internal/tidesofphuket/client/worldTidesClient.go -destination=internal/tools/testing/client/mock_client.go

