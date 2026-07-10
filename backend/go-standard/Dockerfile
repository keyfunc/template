# 构建阶段：准备构建环境
FROM golang:1.25.4-alpine AS builder

# 容器的工作目录
WORKDIR /app

# go mod 代理设置，使用国内的代理加速依赖下载
ENV GOPROXY=https://goproxy.cn,direct

# 复制依赖文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码到容器/app目录
COPY . .

# main 入口开始编译，输出二进制文件 server
RUN go build -o server ./cmd/server/main.go

# 第二阶段：准备一个更小的基础镜像来运行编译好的二进制文件
FROM alpine:3.21

WORKDIR /app

# COPY --from=<阶段名> <源路径> <目标路径>
COPY --from=builder /app/server ./server

# openapi.json 文件拷贝
COPY --from=builder /app/api ./api 

# 启东时执行该二进制程序
CMD ["./server"]
