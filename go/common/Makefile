
all: build test cmd

clean:
	go clean ./...

build:
	go build ./...

test:
	go test -v ./...

cmd:
ifdef CMD
	go install ../cmd/$(CMD)
endif
