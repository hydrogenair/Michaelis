FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/commentcenter-rpc" ]
RUN [ "mkdir", "/usr/local/commentcenter-rpc/etc" ]
WORKDIR /usr/local/commentcenter-rpc
COPY ./commentcenter-rpc /usr/local/commentcenter-rpc
COPY ./etc /usr/local/commentcenter-rpc/etc
CMD ["./commentcenter-rpc", "-f", "etc/comment.yaml"]