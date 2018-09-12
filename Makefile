VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Commit=${COMMIT}"
MODFLAGS=-mod=vendor

BINARY=brewer
PACKAGE=./cmd/brewer

all: dev

clean:
	rm -fr dist/

dev:
	go build ${LDFLAGS} -o dist/${BINARY} ${PACKAGE}

cibuild:
	go build ${MODFLAGS} ${LDFLAGS} -o dist/${BINARY} ${PACKAGE}

dist: darwin linux windows

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${MODFLAGS} ${LDFLAGS} -o dist/${BINARY}-darwin-amd64 ${PACKAGE}

linux:
	GOOS=linux GOARCH=${GOARCH} go build ${MODFLAGS} ${LDFLAGS} -o dist/${BINARY}-linux-amd64 ${PACKAGE}

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${MODFLAGS} ${LDFLAGS} -o dist/${BINARY}-windows-amd64 ${PACKAGE}

test:
	go test ${MODFLAGS} ./...

.PHONY: all clean dev cibuild dist darwin linux windows test
