# 构建阶段
FROM golang:latest AS builder

WORKDIR /gocloudisk

# 复制 go.mod 和 go.sum 并下载依赖
COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download

# 复制源代码
COPY . ./

# 构建二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o clouddisk \
    && chmod +x clouddisk

# 运行阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /gocloudisk/clouddisk /app/clouddisk

# 复制 .env 文件到最终镜像的工作目录
COPY --from=builder /gocloudisk/.env /app/.env

# 确保二进制文件可执行
RUN chmod +x /app/clouddisk

# 设置入口点
ENTRYPOINT ["/app/clouddisk"]