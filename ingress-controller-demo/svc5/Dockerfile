From busybox:glibc

RUN mkdir -p /root/svc5
COPY ./svc5 /root/svc5/svc5
RUN chmod +x /root/svc5/svc5

WORKDIR /root/svc5
ENTRYPOINT ["/root/svc5/svc5"]
