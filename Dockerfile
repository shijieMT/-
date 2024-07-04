# 使用官方的Go镜像作为基础镜像
FROM golang:1.22-alpine

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件44
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制当前目录的所有文件到工作目录
COPY notify-api/ ./notify-api
COPY common/ ./common
COPY notify-model/ ./notify-model


# 编译Go应用
RUN go build -o ./notify-api/notify ./notify-api/notify.go

# 声明服务端口
EXPOSE 8888

WORKDIR /app/notify-api/
# 运行应用
CMD ["./notify"]
