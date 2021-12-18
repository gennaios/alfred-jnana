BINARY_NAME=jnana

GOCMD=go
PROJECTNAME := $(shell basename "$(PWD)")

build:
	GOARCH=amd64 CGO_ENABLED=1 GOOS=darwin $(GOCMD) build -o $(BINARY_NAME) --tags "stat4 foreign_keys vacuum_incr introspect fts5" ./cmd/jnana/*.go