FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/starcenter-api" ]
RUN [ "mkdir", "/usr/local/starcenter-api/etc" ]
WORKDIR /usr/local/starcenter-api
COPY ./starcenter-api /usr/local/starcenter-api
COPY ./etc /usr/local/starcenter-api/etc
CMD ["./starcenter-api", "-f", "etc/starcenter.yaml"]