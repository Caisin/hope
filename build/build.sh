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

function genApiAndClient() {
  for n in "$@"; do
    echo "开始生成模块：$n api"
#    ls "$projectPath/api/$n"
    cd "$projectPath/api/$n" && find . -name "*.proto" -exec  kratos proto client {} \;
  done
  #error_reason.proto
}

#genEnt "${prods[@]}"
genApiAndClient "${prods[@]}"
#for prod in "${prods[@]}"; do
#  #
#  echo "$prod"
#done

sleep 30
