From alpine:latest

COPY ./myhttpserver /root/myhttpserver
RUN chmod +x /root/myhttpserver

WORKDIR /root
ENTRYPOINT ["/root/myhttpserver"]
