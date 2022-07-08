#!/usr/bin/env bash

MODE=$1

if [[ "${MODE}" == "bin" ]]; then
   go version && go env && gcc -v && \
     CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build \
     -ldflags  "-linkmode 'external' -extldflags '-static'" \
     --mod=vendor  -v -o fab-api cmd/main.go
elif [[ "${MODE}" == "docker" ]]; then
    go mod vendor
    # 删除编译镜像
    docker rmi fab-api:v1.0
    docker rmi -f $(docker images | grep "none" | awk '{print $3}')
    #编译镜像并启动
    docker build -t fab-api:v1.0 .
    rm -rf vendor
else
  echo "para is empty bin or docker"
  exit 1
fi
