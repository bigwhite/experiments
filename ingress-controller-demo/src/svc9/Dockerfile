From busybox:glibc

RUN mkdir -p /root/svc9
COPY ./svc9 /root/svc9/svc9
COPY ./server.key /root/svc9/server.key
COPY ./server.crt /root/svc9/server.crt
COPY ./rootCA.pem /root/svc9/rootCA.pem
RUN chmod +x /root/svc9/svc9

WORKDIR /root/svc9
ENTRYPOINT ["/root/svc9/svc9"]
