From busybox:glibc

RUN mkdir -p /root/svcb
COPY ./svcb /root/svcb/svcb
RUN chmod +x /root/svcb/svcb

WORKDIR /root/svcb
ENTRYPOINT ["/root/svcb/svcb"]
