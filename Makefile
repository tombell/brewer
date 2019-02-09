VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT}"
MODFLAGS=-mod=vendor

PLATFORMS:=darwin linux windows

all: dev

clean:
	rm -fr dist/

dev:
	go build ${MODFLAGS} ${LDFLAGS} -o dist/brewer ./cmd/brewer

dist: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$@ GOARCH=amd64 go build ${MODFLAGS} ${LDFLAGS} -o dist/brewer-$@-amd64 ./cmd/brewer

test:
	go test ${MODFLAGS} ./...

.PHONY: all clean dev dist $(PLATFORMS) test
