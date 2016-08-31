.PHONY: all test clean

all:
	$(MAKE) -C erlang

test:
	$(MAKE) -C erlang eunit
	$(MAKE) -C golang test
	$(MAKE) -C test

clean:
	$(MAKE) -C erlang $@
	$(MAKE) -C python $@
	$(MAKE) -C golang $@
	$(MAKE) -C test $@
