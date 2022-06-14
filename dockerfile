# 基础镜像
FROM golang:1.17-alpine
# 复制本地文件到镜像中
ADD . /fiber-mongo-api
# 设置工作目录
WORKDIR /fiber-mongo-api
# 设置环境变量，开启模块支持和设置国内镜像源
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
# 运行命令，通常用于安装应用，安装依赖，配置系统信息（可以有多条 RUN 指令）
RUN go build -o fiber-mongo-api main.go
# 运行命令，通常用于启动程序（只能有一条 CMD 指令）
CMD [ "./fiber-mongo-api" ]