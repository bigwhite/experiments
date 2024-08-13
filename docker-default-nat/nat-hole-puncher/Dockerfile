# 使用 Alpine 作为基础镜像
FROM alpine:latest

# 创建工作目录
WORKDIR /app

# 复制已编译的可执行文件到镜像中
COPY nat-hole-puncher .

# 设置文件权限
RUN chmod +x nat-hole-puncher
