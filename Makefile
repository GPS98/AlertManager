build:
	@go build -o bin/alertmanager

run: build
	@./bin/alertmanager

test:
	@go test -v ./...