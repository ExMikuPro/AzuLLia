FROM golang:1.20.6-alpine3.18 as builder

WORKDIR /usr/src/app

COPY ./ ./

# RUN go env -w GOPROXY=https://goproxy.cn,direct # 国内镜像源

# 设置服务器端口
ENV ServerPort=":80"

# or RunMode="debug"
ENV RunMode=""

# 安装运行以来库
RUN go install && \
    go build -o app .

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /go/bin/GoK-on .
COPY ./staticFile ./staticFile
RUN mkdir "uploadFile"

# 运行容器时执行
CMD ["./GoK-on"]

EXPOSE 80 27017
