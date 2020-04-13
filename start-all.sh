#!/bin/bash

#检查service进程
check_process() {
    sleep 1
    res=`ps aux | grep -v grep | grep "service/bin" | grep $1`
    if [[ $res != '' ]]; then
        echo -e "\033[32m 已启动 \033[0m" "$1"
        return 1
    else
        echo -e "\033[31m 启动失败 \033[0m" "$1"
        return 0
    fi
}

# 编译service可执行文件
build_service() {
    go build -o service/bin/$1 service/$1/main.go
    resbin=`ls service/bin/ | grep $1`
    echo -e "\033[32m 编译完成: \033[0m service/bin/$resbin"
}

# 启动service
run_service() {
    nohup ./service/bin/$1 >> $logpath/$1.log 2>&1 &
    # nohup ./service/bin/$1 --registry=consul >> $logpath/$1.log 2>&1 &
    # nohup go run service/$1/main.go --registry=consul >> $logpath/$1.log 2>&1 &
    sleep 1
    check_process $1
}

# 创建运行日志目录
logpath=/tmp/log/PictureGithubTool
mkdir -p $logpath

# 切换到工程根目录
projectpath=$(cd `dirname $0`; pwd)
cd $projectpath

# 服务列表
services="
apigw
"

# 先停止服务
./stop-all.sh

# 执行编译service
for service in $services
do
    build_service $service
done

# 执行启动service
for service in $services
do
    run_service $service
done

echo '微服务启动完毕.'
