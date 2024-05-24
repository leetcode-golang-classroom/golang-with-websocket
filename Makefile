.PHONY=build

build:
	@go build -o bin/main cmd/main.go cmd/manager.go cmd/client.go cmd/event.go cmd/otp.go

run: build
	@./bin/main

test:
	@go test -v -cover ./test/...