From busybox:glibc

RUN mkdir -p /root/svc6
COPY ./svc6 /root/svc6/svc6
RUN chmod +x /root/svc6/svc6

WORKDIR /root/svc6
ENTRYPOINT ["/root/svc6/svc6"]
