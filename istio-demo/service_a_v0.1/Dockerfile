From busybox:glibc

RUN mkdir -p /root/svca
COPY ./svca /root/svca/svca
RUN chmod +x /root/svca/svca

WORKDIR /root/svca
ENTRYPOINT ["/root/svca/svca"]
