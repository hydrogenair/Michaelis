FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/postcenter-rpc" ]
RUN [ "mkdir", "/usr/local/postcenter-rpc/etc" ]
WORKDIR /usr/local/postcenter-rpc
COPY ./postcenter-rpc /usr/local/postcenter-rpc
COPY ./etc /usr/local/postcenter-rpc/etc
CMD ["./postcenter-rpc", "-f", "etc/postcenter.yaml"]