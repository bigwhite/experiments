From busybox:glibc

RUN mkdir -p /root/svc8
COPY ./svc8 /root/svc8/svc8
COPY ./key.pem /root/svc8/key.pem
COPY ./cert.pem /root/svc8/cert.pem
RUN chmod +x /root/svc8/svc8

WORKDIR /root/svc8
ENTRYPOINT ["/root/svc8/svc8"]
