FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/connectioncenter-api" ]
RUN [ "mkdir", "/usr/local/connectioncenter-api/etc" ]
WORKDIR /usr/local/connectioncenter-api
COPY ./connectioncenter-api /usr/local/connectioncenter-api
COPY ./etc /usr/local/connectioncenter-api/etc
CMD ["./connectioncenter-api", "-f", "etc/connectioncenter.yaml"]