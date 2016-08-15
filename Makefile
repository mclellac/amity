NO_COLOR=$(shell echo  "\033[0m")
OK_COLOR=$(shell echo  "\033[32;01m")
ERROR_COLOR=$(shell echo  "\033[31;01m")
WARN_COLOR=$(shell echo  "\033[33;01m")
SOURCE=$(go list ./... | grep -v "tests")
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOFMT=gofmt -w
DEPS=$(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
PACKAGES = $(shell go list ./...)

default: build

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

updatedeps:
	@echo "$(OK_COLOR)==> Updating all dependencies$(NO_COLOR)"
	@go get -d -u ./...
	@echo $(DEPS) | xargs -n1 go get -d -u

format:
	@echo "$(OK_COLOR)==> Formatting$(NO_COLOR)"
	$(GOFMT) $(PACKAGES)

build:
	@echo "$(OK_COLOR)==> Building$(NO_COLOR)"
	$(GOBUILD) -o ./amityd ./cmd/amityd/
	$(GOBUILD) -o ./amity  ./cmd/amity/

clean:
	go clean -i -r -x
	rm ./amityd && rm ./amity

#migrate:
#	./amityd --config amity.gcfg migratedb

install:
	@echo "$(OK_COLOR)==> Installing$(NO_COLOR)"
	go install ./cmd/amityd
	go install ./cmd/amity

lint:
	@echo "$(OK_COLOR)==> Linting$(NO_COLOR)"
	${GOPATH}/bin/golint .

vet:
	go vet ./cmd/amityd/
	go vet ./cmd/amity/
	go vet ./lib/api/
	go vet ./libamityd/
	go vet ./lib/amity/

test:
	./amityd --config amityd.gcfg start & pid=$$!; cd tests && go test; kill $$pid

all: format lint test
