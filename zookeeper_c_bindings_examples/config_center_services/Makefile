all:
	gcc -g -std=gnu99 -o ccs_agent ccs_agent.c -DTHREADED -I/usr/local/include/zookeeper -lzookeeper_mt -lpthread
	gcc -g -std=gnu99 -o app app.c -DTHREADED -I/usr/local/include/zookeeper -lzookeeper_mt -lpthread

clean:
	rm ccs_agent app
