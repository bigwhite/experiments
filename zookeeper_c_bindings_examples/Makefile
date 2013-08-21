all:
	gcc -g -std=gnu99 -o election election.c -DTHREADED -I/usr/local/include/zookeeper -lzookeeper_mt -lpthread

clean:
	rm election
