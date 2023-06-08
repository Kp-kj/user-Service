# 基于Go官方镜像作为基础镜像
FROM golang:1.20-alpine

# 设置工作目录
WORKDIR /app

# 将项目文件复制到镜像中
COPY . .

# 构建项目
RUN go build -o user-server

# 暴露端口
EXPOSE 8080

# 运行项目
CMD ["./user-server"]