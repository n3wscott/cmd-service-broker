ifeq ($(GOBIN),)
GOBIN := $(GOPATH)/bin
endif

all: build test verify

BINDIR        ?= bin
TOP_SRC_DIRS   = cmd pkg
SRC_DIRS       = $(shell sh -c "find $(TOP_SRC_DIRS) -name \\*.go \
                   -exec dirname {} \\; | sort | uniq")
NEWEST_GO_FILE = $(shell find $(SRC_DIRS) -name \*.go -exec $(STAT) {} \; \
                   | sort -r | head -n 1 | sed "s/.* //")


run: build
	./bin/cmd-broker --v 1 -logtostderr

clean:
	go clean

dependencies:
	go get -u github.com/golang/lint/golint
	go get -u github.com/kardianos/govendor
	$(GOBIN)/govendor sync

build: cmd-broker

cmd-broker: $(BINDIR)/cmd-broker
$(BINDIR)/cmd-broker: cmd/cmd-broker $(NEWEST_GO_FILE)
	go build -o $@ ./cmd/cmd-broker

verify:
	go vet `go list ./... | grep -v /vendor/`
	golint `go list ./... | grep -v /vendor/`

test: check
	go test `go list ./... | grep -v /vendor/`

ci: clean dependencies build test

default: build

.PHONY: test
