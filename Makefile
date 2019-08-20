GOFMT=gofmt -w
DEPS=$(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
PACKAGES := $(shell go list ./...)

# Prettify output
UNAME_S	 := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	RESET = $(shell echo -e "\033[0m")
	GREEN = $(shell echo -e "\033[32;01m")
	ERROR = $(shell echo -e "\033[31;01m")
	WARN  = $(shell echo -e "\033[33;01m")
endif
ifeq ($(UNAME_S),Darwin)
	RESET := $(shell echo "\033[0m")
	GREEN := $(shell echo "\033[32;01m")
	ERROR := $(shell echo "\033[31;01m")
	WARN  := $(shell echo "\033[33;01m")
endif

#default: build

deps:
	@echo "$(GREEN)==> Installing dependencies$(RESET)"
	@go get -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

updatedeps:
	@echo "$(GREEN)==> Updating all dependencies$(RESET)"
	@go get -d -u ./...
	@echo $(DEPS) | xargs -n1 go get -d -u

format:
	@echo "$(GREEN)==> Formatting$(RESET)"
	$(foreach ENTRY,$(PACKAGES),$(GOFMT) $(GOPATH)/src/$(ENTRY);)

build:
	@echo "$(GREEN)==> Building$(RESET)"
	go build -o ./bin/amityd ./cmd/amityd/
	go build -o ./bin/amity  ./cmd/amity/

clean:
	go clean -i -r -x
	rm ./bin/*

migrate:
	./bin/amityd --config amityd.conf migratedb

install:
	@echo "$(GREEN)==> Installing$(RESET)"
	go install ./cmd/amityd
	go install ./cmd/amity

lint:
	@echo "$(GREEN)==> Linting$(RESET)"
	${GOPATH}/bin/golint .

vet:
	go vet ./cmd/amityd/
	go vet ./cmd/amity/
	go vet ./lib/api/
	#go vet ./lib/amityd/
	#go vet ./lib/amity/

test:
	./bin/amityd --config amityd.conf start & pid=$$!; cd tests && go test; kill $$pid


all: format lint vet build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
