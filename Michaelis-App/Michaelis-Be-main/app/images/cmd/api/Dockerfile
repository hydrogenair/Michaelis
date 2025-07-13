FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/imagecenter-api" ]
RUN [ "mkdir", "/usr/local/imagecenter-api/etc" ]
WORKDIR /usr/local/imagecenter-api
COPY ./imagecenter-api /usr/local/imagecenter-api
COPY ./etc /usr/local/imagecenter-api/etc
CMD ["./imagecenter-api", "-f", "etc/imagecenter.yaml"]