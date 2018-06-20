From busybox:glibc

RUN mkdir -p /root/svc3
COPY ./svc3 /root/svc3/svc3
RUN chmod +x /root/svc3/svc3

WORKDIR /root/svc3
ENTRYPOINT ["/root/svc3/svc3"]
