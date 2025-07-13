FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/usercenter-api" ]
RUN [ "mkdir", "/usr/local/usercenter-api/etc" ]
WORKDIR /usr/local/usercenter-api
COPY ./usercenter-api /usr/local/usercenter-api
COPY ./etc /usr/local/usercenter-api/etc
CMD ["./usercenter-api", "-f", "etc/usercenter.yaml"]