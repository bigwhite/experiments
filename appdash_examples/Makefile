
SUBDIRS = appdash frontservice backendservices

all:
	@for i in $(SUBDIRS); do \
		echo "\n== Enter '$$i' =="; \
		(cd $$i &&make -f Makefile) || exit 1; \
	done

clean:
	@for i in $(SUBDIRS); do \
		echo "\n== Enter '$$i' =="; \
		(cd $$i &&make clean -f Makefile) || exit 1; \
	done
