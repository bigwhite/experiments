From busybox:glibc

RUN mkdir -p /root/svc1
COPY ./svc1 /root/svc1/svc1
RUN chmod +x /root/svc1/svc1

WORKDIR /root/svc1
ENTRYPOINT ["/root/svc1/svc1"]
