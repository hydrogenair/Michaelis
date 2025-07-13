FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/postcenter-api" ]
RUN [ "mkdir", "/usr/local/postcenter-api/etc" ]
WORKDIR /usr/local/postcenter-api
COPY ./postcenter-api /usr/local/postcenter-api
COPY ./etc /usr/local/postcenter-api/etc
CMD ["./postcenter-api", "-f", "etc/postcenter.yaml"]