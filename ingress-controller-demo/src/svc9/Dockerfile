From busybox:glibc

RUN mkdir -p /root/svc9
COPY ./svc9 /root/svc9/svc9
COPY ./key.pem /root/svc9/key.pem
COPY ./cert.pem /root/svc9/cert.pem
RUN chmod +x /root/svc9/svc9

WORKDIR /root/svc9
ENTRYPOINT ["/root/svc9/svc9"]
