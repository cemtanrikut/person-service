# Makefile for the Person Service

BINARY_NAME=person-service

# Go komutları
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

all: build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
deps:
	$(GOGET) -u ./...
run:
	$(GORUN) main.go
tidy:
	$(GOCMD) mod tidy
info:
	@echo "Binary Name: $(BINARY_NAME)"
	@echo "Dependencies:"
	@$(GOCMD) list -m all

.PHONY: build test clean deps run tidy info