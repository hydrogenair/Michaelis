FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/chatcenter-api" ]
RUN [ "mkdir", "/usr/local/chatcenter-api/etc" ]
WORKDIR /usr/local/chatcenter-api
COPY ./chatcenter-api /usr/local/chatcenter-api
COPY ./etc /usr/local/chatcenter-api/etc
CMD ["./chatcenter-api", "-f", "etc/chat.yaml"]