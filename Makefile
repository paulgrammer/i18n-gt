.PHONY: build dev build-linux

BINARY_NAME = i18n-gt
BUILD_DIR = $(PWD)/build
BINARY_PATH = $(BUILD_DIR)/$(BINARY_NAME)

build:
	go build -o $(BINARY_PATH) main.go

dev: 
	export GOOGLE_APPLICATION_CREDENTIALS=account.json && go run main.go --input=input.json --output=output.json --source=en --target=pl

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_PATH)-linux-amd64 main.go