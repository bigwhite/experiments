
LIBMEMCACHED_INCLUDE = /home1/tonybai/temp/libmemcached-1.0.17/dist/include/
LIBMEMCACHED_LIB = /home1/tonybai/temp/libmemcached-1.0.17/dist/lib

all:
	gcc -o pack_sys -g pack_sys.c -I$(LIBMEMCACHED_INCLUDE) -L $(LIBMEMCACHED_LIB) -lmemcached -lstdc++
clean:
	rm -f core pack_sys
