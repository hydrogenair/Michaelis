FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/imagecenter-rpc" ]
RUN [ "mkdir", "/usr/local/imagecenter-rpc/etc" ]
WORKDIR /usr/local/imagecenter-rpc
COPY ./imagecenter-rpc /usr/local/imagecenter-rpc
COPY ./etc /usr/local/imagecenter-rpc/etc
CMD ["./imagecenter-rpc", "-f", "etc/imagecenter.yaml"]