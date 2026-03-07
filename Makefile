VERSION ?= dev

.PHONY: build run clean fmt lint test test-docker

build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/switchboard ./cmd/switchboard

run: build
	./bin/switchboard

clean:
	rm -rf bin/

fmt:
	gofumpt -w .

lint:
	golangci-lint run ./...

test:
	go test ./... -v

test-docker:
	@command -v docker >/dev/null 2>&1 || { echo "Error: Docker is required but not found."; exit 1; }
	@docker info >/dev/null 2>&1 || { echo "Error: Docker daemon is not running."; exit 1; }
	@mkdir -p test-results
	docker compose -f docker-compose.test.yml run --rm test
