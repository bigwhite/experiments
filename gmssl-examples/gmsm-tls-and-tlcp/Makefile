
all:
	go build -o echoserver gmsmdemo/server
	go build -o echoserver_gm gmsmdemo/server_gm
	go build -o echo_tls_client gmsmdemo/tlsclient
	go build -o echo_tlcp_client gmsmdemo/tlcpclient
	go build -o echo_tls_client_gm gmsmdemo/tlsclient_gm

clean:
	rm -fr ./echoserver ./echoserver_gm ./echo_tls_client ./echo_tls_client_gm ./echo_tlcp_client
