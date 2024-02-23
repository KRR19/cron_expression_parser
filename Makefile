GOCMD=go
SRCMAIN=./cmd
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=cron_expression_parser

build:  
	$(GOBUILD) -o ./bin/ -v ./cmd/...
test:  
	$(GOTEST) -v ./...
clean:  
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run: 
	$(GOCMD) run $(SRCMAIN)/main.go
deps:
	$(GOGET) ./...