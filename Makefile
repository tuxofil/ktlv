.PHONY: all test clean

all:
	$(MAKE) -C erlang

test:
	$(MAKE) -C erlang eunit
	$(MAKE) -C test

clean:
	$(MAKE) -C erlang $@
	$(MAKE) -C python $@
	$(MAKE) -C test $@
