fmt:
	@echo "\nFormating\n"
	@go fmt ./...

test: fmt
	@echo "\nRunning tests\n"
	@echo go test ./..

build:
	@echo "\nRuning Go build\n"
	@go build

up: stop
	@echo "\nStarting local database"
	@docker-compose up -d

run: up
	@echo "\nRunning Go run\n"
	@go run main.go 

stop:
	@echo "\nStop Container\n"
	@docker-compose stop
