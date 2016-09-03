TESTS = erlang-python erlang-golang python

.PHONY: all clean $(TESTS)

all: $(TESTS)

clean: $(TESTS)

$(TESTS):
	$(MAKE) -C $@ $(MAKECMDGOALS)
