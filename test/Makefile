TESTS = erlang-golang erlang-python python

.PHONY: all clean $(TESTS)

all: $(TESTS)

clean: $(TESTS)

$(TESTS):
	$(MAKE) -C $@ $(MAKECMDGOALS)
