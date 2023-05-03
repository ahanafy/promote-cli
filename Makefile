BINARY_NAME=promote-cli
# get and set GOOS automatically
ifeq ($(OS),Windows_NT)
	GOOS=windows
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		GOOS=linux
	endif
	ifeq ($(UNAME_S),Darwin)
		GOOS=darwin
	endif
endif

.PHONY: all
all: dep vet lint test run

.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows main.go

.PHONY: run
run: build
	"./bin/${BINARY_NAME}-${GOOS}" --config promote-cli.yaml --check development

.PHONY: clean
clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

.PHONY: test
test:
	go test ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: dep
dep:
	go mod download

.PHONY: vet
vet:
	go vet

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lintception
lintception:
	go run github.com/mrtazz/checkmake/cmd/checkmake@latest Makefile
	golangci-lint run --fix