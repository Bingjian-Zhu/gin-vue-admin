# 构建：使用golang:1.13版本
FROM golang:1.13 as build

# 容器环境变量添加，会覆盖默认的变量值
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# # 设置工作区
# WORKDIR /go/cache

# # 把go.mod和go.sum添加到/go/cache目录
# ADD go.mod .
# ADD go.sum .
# # 执行go mod download
# RUN go mod download

# 设置工作区
WORKDIR /go/release

# 把全部文件添加到/go/release目录
ADD . .

# 编译：把cmd/main.go编译成可执行的二进制文件，命名为app
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app cmd/main.go

# 运行：使用scratch作为基础镜像
FROM scratch as prod

# 在build阶段复制时区到
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 在build阶段复制可执行的go二进制文件app
COPY --from=build /go/release/app /
# 在build阶段复制配置文件
COPY --from=build /go/release/config ./config

# 启动服务
ENTRYPOINT ["/app"]