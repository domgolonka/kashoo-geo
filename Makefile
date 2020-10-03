GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=kashoo

all: clean swag-init build swag-code
.PHONY: build
build:
	@echo "building app...."
	$(GOBUILD) -o $(BINARY_NAME) $(shell pwd)/cmd/main.go
.PHONY: run build
run: build
	@echo "Running application...."
	./$(BINARY_NAME)
.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
.PHONY: swagger-init
swag-init:
	@echo "Starting Swag Docs...."
	swag init -g "./cmd/main.go"
.PHONY: swagger-code
swag-code:
	@echo "Starting Codegen..."
	swagger-codegen generate -l go -i ./docs/swagger.json -o client
