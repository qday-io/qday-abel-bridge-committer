# 构建阶段
FROM golang:1.22 as builder

# 安装依赖和工具
RUN apt-get update && apt-get install -y protobuf-compiler git && \
    cd /tmp && git clone https://github.com/googleapis/googleapis.git && \
    cp -r /tmp/googleapis/* /usr/local/include/ && \
    go install github.com/golang/protobuf/protoc-gen-go@latest

# 设置工作目录
WORKDIR /src

# 拷贝代码
COPY . .

# 下载依赖
RUN go mod download

# （如需生成 proto 文件，取消下一行注释）
# RUN protoc --proto_path=proto --go_out=pkg proto/*.proto

# 编译
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/committer main.go

# 运行阶段
FROM alpine:latest

WORKDIR /app

# 拷贝二进制文件
COPY --from=builder /app/committer /app/committer

# 启动命令
CMD ["/app/committer"]

