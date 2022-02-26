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

test: ## Test the application localy.
	@echo "Testing application:"
	@CGO_ENABLED=1 go test ./... -race -count=1 -cover -coverprofile=coverage.txt && go tool cover -func=coverage.txt | tail -n1 | awk '{print "Total test coverage: " $$3}'
	@rm coverage.txt