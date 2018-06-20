From busybox:glibc

RUN mkdir -p /root/svc2
COPY ./svc2 /root/svc2/svc2
COPY ./key.pem /root/svc2/key.pem
COPY ./cert.pem /root/svc2/cert.pem
RUN chmod +x /root/svc2/svc2

WORKDIR /root/svc2
ENTRYPOINT ["/root/svc2/svc2"]
