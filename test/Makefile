TESTS = erlang-python

.PHONY: all clean $(TESTS)

all: $(TESTS)

clean: $(TESTS)

$(TESTS):
	$(MAKE) -C $@ $(MAKECMDGOALS)
