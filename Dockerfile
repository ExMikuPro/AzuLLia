FROM golang:1.20.6-alpine3.18 as builder

WORKDIR /usr/src/app

COPY ./ ./

# 国内镜像源
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 安装运行依赖库
RUN go install && \
    go build -o app .

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /usr/src/app/app .

COPY ./staticFile ./staticFile

COPY ./Service.env ./Service.env

RUN mkdir -p "uploadFile"

# 运行容器时执行
CMD ["./app"]

EXPOSE 80 27017
