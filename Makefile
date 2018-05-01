.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY: run exec build clean deps test

run: build exec clean

exec:
	./bin/app

build:
	go build -o bin/app -ldflags '-s -w -extldflags "-static"' main.go

clean:
	rm -rf bin

deps:
	go get -d -u -v ./...

test:
	GOCACHE=off go test -failfast -v ./pkg/...
