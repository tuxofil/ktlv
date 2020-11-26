.PHONY: all clean

ifndef ($(HOME))
    export HOME = /tmp
endif

all: clean
	../erlang-python/erlang-write.escript
	GOPATH=`pwd`/../../golang go run golang-read.go

clean:
	rm -f -- object.bin ktlv.hrl
