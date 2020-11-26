.PHONY: all test cover clean

ifndef ($(HOME))
   export HOME = /tmp
endif

all:

test:
	cd src/ktlv && go test -v -race -coverprofile=.cover.out ./...

cover: test
	cd src/ktlv && go tool cover -html=.cover.out

clean:
	rm -f src/ktlv/.cover.out
