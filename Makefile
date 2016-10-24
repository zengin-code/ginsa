PREFIX:=$(shell pwd)
BINDIR=$(PREFIX)/bin
BINARY=$(BINDIR)/$(shell basename `pwd`)

SRC=$(shell find . -name "*.go")

VERSION=$(shell cat VERSION)
REVISION=$(shell git rev-parse --short HEAD)
LDFLAGS=-X github.com/zengin-code/ginsa.VERSION=$(VERSION) \
				-X github.com/zengin-code/ginsa.REVISION=$(REVISION)
GOFLAG=-ldflags "$(LDFLAGS)"

all: build

$(BINARY): $(SRC)
	@go build $(GOFLAG) -o $@ github.com/zengin-code/ginsa/cli

build: $(BINARY)

deps:
	@git grep github.com | sed -e "s/\"$$//" | sed -e "s/^.*\"//" | sort | uniq | grep -v depscribe