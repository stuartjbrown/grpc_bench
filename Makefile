.PHONY: all gen build install test lint get publish
PACKAGE=github.com/stuartjbrown/grpc_bench

all: get install

gen:
	go generate .

build:
	go build -a -o bin/grpc_bench ${PACKAGE}

test: lint
	go test ${PACKAGE}

cover: test
	go tool cover -html=coverage.out

lint:
	go vet ${PACKAGE}
	golint ${PACKAGE}

get:
	go get -u github.com/golang/dep/cmd/dep

bench:
	go test ./... -bench=Bench
