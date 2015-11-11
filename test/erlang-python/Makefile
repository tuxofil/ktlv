.PHONY: all clean

all: clean
	ln -sf ../../erlang/include/ktlv.hrl .
	./erlang-write.escript
	PYTHONPATH=../../python ./python-read.py

clean:
	rm -f -- object.bin *.pyc *.pyo ktlv.hrl
