APP = ktlv

VERSION = $(shell head -1 ../debian/changelog | awk '{print $$2}' | sed -r 's/[()]//g')

.PHONY: all compile html clean eunit dialyze

COPTS := -v -Werror

ifdef TEST
COPTS := $(COPTS) -DTEST
endif

ifdef DEBUG
COPTS := $(COPTS) +debug_info
endif

all: html compile

compile:
	mkdir -p ebin
	sed "s/{{VERSION}}/$(VERSION)/" src/$(APP).app.in > ebin/$(APP).app
	erlc $(COPTS) -Iinclude -o ebin src/*.erl

EDOC_OPTS = {application,$(APP)}, {preprocess,true}
html:
	sed "s/{{VERSION}}/$(VERSION)/" doc/overview.edoc.in > doc/overview.edoc
	erl -noinput -eval 'edoc:application($(APP),".",[$(EDOC_OPTS)]),halt()'

eunit:
	$(MAKE) TEST=y clean compile
	erl -noinput -pa ebin \
		-eval 'ok=eunit:test({application,$(APP)},[verbose]),halt()'

PLT = .dialyzer_plt
DIALYZER_OPTS = -Wunmatched_returns -Werror_handling
DIALYZER_APPS = kernel stdlib

dialyze: $(PLT)
	$(MAKE) DEBUG=y clean compile
	dialyzer --plt $< -r . $(DIALYZER_OPTS) --src
	dialyzer --plt $< -r . $(DIALYZER_OPTS)

$(PLT):
	dialyzer --build_plt --output_plt $@ --apps $(DIALYZER_APPS)

all-tests:
	$(MAKE) eunit
	$(MAKE) dialyze

clean:
	rm -rf -- ebin doc/*.html doc/*.css doc/*.png doc/edoc-info \
	    erl_crash.dump doc/overview.edoc
