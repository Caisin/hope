#!/usr/bin/env bash
# 项目路径
projectPath="D:/work/code/go/hope"
echo "项目路径为：$projectPath"
# 模块
prods=({"admin","param","novel"})
echo 模块为: "${prods[@]}"

#函数定义

##生成ent
function genEnt() {
  for n in "$@"; do
    echo "开始生成模块：$n ent"
    cd "$projectPath/apps/$n/internal/data" && ent generate ./ent/schema
  done
}

#生成api
function genApiAndClient() {
  for n in "$@"; do
    echo "开始生成模块：$n api"
    #    ls "$projectPath/api/$n"
    cd "$projectPath/api/$n" && find . -name "*.proto" -exec kratos proto client {} \;
    echo "开始生成模块：$n server"
    cd "$projectPath/api/$n" && find . -name "*.proto" -exec kratos proto server {} -t "$projectPath/apps/$n/internal/service" \;
    echo "开始生成模块：$n wire"
    cd "$projectPath/apps/$n/cmd/$n" && wire
  done
  #error_reason.proto
}
#生成配置
function genConfig() {
  for n in "$@"; do
    echo "开始生成模块：$n Config"
    cd "$projectPath/apps/$n/internal" && find . -name "*.proto" -exec kratos proto {} \;
  done
}
#配置
genConfig "${prods[@]}"
#ent生成
genEnt "${prods[@]}"
# api生成
genApiAndClient "${prods[@]}"
#for prod in "${prods[@]}"; do
#  #
#  echo "$prod"
#done

sleep 1m
