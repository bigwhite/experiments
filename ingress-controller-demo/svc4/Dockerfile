From busybox:glibc

RUN mkdir -p /root/svc4
COPY ./svc4 /root/svc4/svc4
RUN chmod +x /root/svc4/svc4

WORKDIR /root/svc4
ENTRYPOINT ["/root/svc4/svc4"]
