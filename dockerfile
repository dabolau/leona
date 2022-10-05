# 基础镜像
FROM golang:1.18-alpine
# 复制本地文件到镜像中
ADD . /leona
# 设置工作目录
WORKDIR /leona
# 运行命令，通常用于启动程序（只能有一条 CMD 指令）
CMD [ "./leona" ]