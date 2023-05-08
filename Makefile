BINARY_NAME := i18n-gt
BUILD_DIR := $(PWD)/build
BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME)
VERSION := 1.0.0
BUILD_TIME := $(shell date +%FT%T%z)
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

.PHONY: build clean dev

dev: 
	export GOOGLE_APPLICATION_CREDENTIALS=account.json && go run main.go --input=input.json --output=output.json --source=en --target=pl

build: build-linux-amd64 build-linux-arm build-darwin-amd64 build-windows-amd64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_PATH)-$(VERSION)-linux-amd64

build-linux-arm:
	GOOS=linux GOARCH=arm go build $(LDFLAGS) -o $(BINARY_PATH)-$(VERSION)-linux-arm

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_PATH)-$(VERSION)-darwin-amd64

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_PATH)-$(VERSION)-windows-amd64.exe

clean:
	rm -f $(BINARY_PATH)-*
