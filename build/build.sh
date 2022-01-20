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
    cd "$projectPath/apps/$n/internal/data" && ent generate --idtype=int64 ./ent/schema
  done
}

#生成api
function genApi() {
  for n in "$@"; do
    echo "开始生成模块：$n api"
    #    ls "$projectPath/api/$n"
    cd "$projectPath/api/$n" && find . -name "*.proto" -exec kratos proto client --proto_path=$projectPath/third_party {} \;
  done
  #error_reason.proto
}

#生成api
function genWire() {
  for n in "$@"; do
    echo "开始生成模块：$n wire"
    cd "$projectPath/apps/$n/cmd/$n" && wire
  done
  #error_reason.proto
}
#生成配置
function genConfig() {
  for n in "$@"; do
    echo "开始生成模块：$n Config"
    cd "$projectPath/apps/$n/internal" && find . -name "*.proto" -exec kratos proto client {} \;
  done
}
mode=$1
case $mode in
wire)
  echo 'You select wire'
  genWire "${prods[@]}"
  ;;
ent)
  echo 'You select ent'
  genEnt "${prods[@]}"
  ;;
api)
  echo 'You select api'
  genApi "${prods[@]}"
  ;;
config)
  echo 'You select config'
  genConfig "${prods[@]}"
  ;;
all)
  echo 'You select all'
  genWire "${prods[@]}"
  genEnt "${prods[@]}"
  genApi "${prods[@]}"
  genConfig "${prods[@]}"
  ;;
*)
  "Usage: $mode [wire|ent|api|config|all]"
  sleep 1m
  exit 1
  ;;
esac
sleep 1m

