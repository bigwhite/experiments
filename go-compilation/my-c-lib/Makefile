.PHONY:  all static

all:
	gcc -c -fPIC -o mylib.o mylib.c
	gcc -shared -o libmylib.so mylib.o
static:
	gcc -c -fPIC -o mylib.o mylib.c
	ar rcs libmylib.a mylib.o
