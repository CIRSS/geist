include ./common/Makefile

package:
	go build -o ./.build ./... 

clean:
	go clean ./...
	rm -rf ./.build/*
	chmod -R +w ${GOPATH}
	rm -rf ${GOPATH}/*
