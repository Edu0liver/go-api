.PHONY: default run build test docs clean
# Variables
APP_NAME=usersmanager

# Tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init -g internal/cmd/main.go
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs