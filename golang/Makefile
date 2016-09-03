.PHONY: all test fmt clean

all:

test:
	GOPATH=`pwd` go test -v ktlv

fmt:
	find . -type f -name \*.go -exec gofmt -w '{}' ';'

clean:
