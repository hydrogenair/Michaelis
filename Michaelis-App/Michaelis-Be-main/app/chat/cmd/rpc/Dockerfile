FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/chatcenter-rpc" ]
RUN [ "mkdir", "/usr/local/chatcenter-rpc/etc" ]
WORKDIR /usr/local/chatcenter-rpc
COPY ./chatcenter-rpc /usr/local/chatcenter-rpc
COPY ./etc /usr/local/chatcenter-rpc/etc
CMD ["./chatcenter-rpc", "-f", "etc/chat.yaml"]