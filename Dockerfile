FROM golang:1.20.6

WORKDIR /usr/src/app

COPY ./ ./

# RUN go env -w GOPROXY=https://goproxy.cn,direct # 国内镜像源

# 安装运行以来库
RUN go install && \
    go build -o app .

# 设置服务器端口
ENV ServerPort=":80"

# or RunMode="debug"
ENV RunMode=""

# 运行容器时执行
CMD ["./app"]

EXPOSE 80 27017
