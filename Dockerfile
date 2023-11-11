FROM golang:1.20.6-alpine3.18 as builder

WORKDIR /usr/src/app

COPY ./ ./

RUN go env -w GOPROXY=https://goproxy.cn,direct # 国内镜像源

# 安装运行依赖库
RUN go install && \
    go build -o app .

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /usr/src/app/app .

COPY ./staticFile ./staticFile

RUN mkdir -p "uploadFile"

# 设置服务器端口
ENV ServerPort=":80"

# or RunMode="debug"
ENV RunMode="release"

# 运行容器时执行
CMD ["./app"]

EXPOSE 80 27017
