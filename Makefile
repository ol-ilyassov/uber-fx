
BINARY_NAME=app
BINARY_PATH=cmd/build
APP_PATH=cmd

.SILENT:

.PHONY: run
run:
	clear && go run $(APP_PATH)/main.go

.PHONY: build
build:
	mkdir -p $(BINARY_PATH)
	CGO_ENABLED=0 go build -o $(BINARY_PATH)/$(BINARY_NAME) $(APP_PATH)/main.go
