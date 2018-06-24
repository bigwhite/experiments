From busybox:glibc

RUN mkdir -p /root/svc7
COPY ./svc7 /root/svc7/svc7
RUN chmod +x /root/svc7/svc7

WORKDIR /root/svc7
ENTRYPOINT ["/root/svc7/svc7"]
