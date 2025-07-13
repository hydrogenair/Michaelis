## 设置基础镜像
FROM ubuntu:20.04
#
#RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY="https://goproxy.cn,direct"
# 创建工作目录
WORKDIR /app


RUN ["mkdir",""]
# 复制源代码到容器中
COPY . app


# 设置容器启动命令
CMD ["./main"]
