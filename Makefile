GO := go
GOCMD := $(GO)
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOVET := $(GOCMD) vet
GOFMT := $(GOCMD) fmt
GORUN := $(GOCMD) run

BINARY_NAME := go-microservice

default: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	$(GOTEST) ./...

deps:
	$(GOGET) -v

vet:
	$(GOVET) ./...

lint:
	$(GOFMT) ./...

dev:
	$(GORUN) main.go

.PHONY: build clean test deps vet fmt

