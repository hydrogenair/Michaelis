FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/starcenter-rpc" ]
RUN [ "mkdir", "/usr/local/starcenter-rpc/etc" ]
WORKDIR /usr/local/starcenter-rpc
COPY ./starcenter-rpc /usr/local/starcenter-rpc
COPY ./etc /usr/local/starcenter-rpc/etc
CMD ["./starcenter-rpc", "-f", "etc/starcenter.yaml"]