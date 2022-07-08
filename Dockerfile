FROM golang:1.16 as builder
ARG Version
ARG CommitVersion
ARG BuildTime
LABEL version=$Version comshbuimit=$CommitVersion create_time=$BuildTime

ADD . /fab
WORKDIR /fab
RUN  go version && go env && gcc -v && \
     CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build \
     --mod=vendor  -v -o fab-api cmd/main.go

#生成中间镜像后,将build之后的可执行文件考到新的镜像中
FROM golang:1.14.4-alpine3.12 as fab-api
ARG Version
ARG CommitVersion
ARG BuildTime
LABEL version=$Version commit=$CommitVersion create_time=$BuildTime
COPY --from=builder  /fab/fab-api /usr/local/bin
COPY --from=builder  /fab/deployments/config/fab-config.yaml /etc/fab/
COPY --from=builder  /fab/deployments/config/network.yaml /etc/fab/
# 切换软件源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
WORKDIR /fab/fab-api
#容器内部开放端口
ENTRYPOINT ["fab-api"]
