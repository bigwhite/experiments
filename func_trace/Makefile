#
# Makefile for func_trace share lib
#

OS := $(shell uname -s)

LIBS = 

ifeq ($(OS), SunOS)
	DEFS = -D_SOLARIS_TRACE
	LIBS += -lfunc_trace
else ifeq ($(OS), Linux)
	DEFS = -D_LINUX_TRACE
else
	$(error $(OS) is not supported!)
endif


all:
	gcc $(DEFS) -fPIC -shared -o libfunc_trace.so func_trace.c 
	gcc -g example.c -o example -finstrument-functions $(LIBS)

clean:
	rm -fr example libfunc_trace.so
