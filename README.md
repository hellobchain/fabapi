# fabapi
基于go语言的fabric sdk API server  支持2.X
# 编译
````
# 镜像编译
sh build/build.sh docker
# 二进制文件编译
sh build/build.sh bin
````
# 配置
````
# 环境变量
export FAB_NETWORK_PATH="deployments/config/fab-config.yaml"     # fabconfig.yaml的路径
expoer FAB_CONFIG_PATH="deployments/config/network.yaml"      # network.yaml的路径
# 修改fabconfig.yaml
deployments/config/fab-config.yaml
# 修改network.yaml
deployments/config/network.yaml
````

# 部署
````
# 运行
sh deployments/docker/start.sh
# 停止
sh deployments/docker/stop.sh

# 二进制文件运行
sh ./fab-server 
````
