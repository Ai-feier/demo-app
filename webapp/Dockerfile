# 第一阶段：构建应用程序
FROM golang:1.20 AS builder

# 设置工作目录
WORKDIR /app

# 复制所有文件并下载依赖
COPY . .
RUN go mod tidy

# 编译应用程序
RUN go build -o main .

# 第二阶段：创建一个更小的运行时镜像
FROM alpine:latest

# 安装依赖
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

# 运行应用程序
CMD ["./main"]
