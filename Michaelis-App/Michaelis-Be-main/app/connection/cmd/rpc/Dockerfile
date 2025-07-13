FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/connectioncenter-rpc" ]
RUN [ "mkdir", "/usr/local/connectioncenter-rpc/etc" ]
WORKDIR /usr/local/connectioncenter-rpc
COPY ./connectioncenter-rpc /usr/local/connectioncenter-rpc
COPY ./etc /usr/local/connectioncenter-rpc/etc
CMD ["./connectioncenter-rpc", "-f", "etc/connectioncenter.yaml"]