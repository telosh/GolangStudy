#!/bin/bash

# 引数チェック
if [ $# -eq 0 ]; then
    echo "Usage: ./start.sh <app_directory>"
    echo "Example: ./start.sh helloworld"
    exit 1
fi

APP_DIR=$1

# アプリケーションディレクトリの存在確認
if [ ! -d "$APP_DIR" ]; then
    echo "Error: Directory '$APP_DIR' does not exist"
    exit 1
fi

# docker-compose.yamlの一時的なコピーを作成
cp docker/docker-compose.yaml docker/docker-compose.tmp.yaml

# サービス名を動的に設定
sed -i "s/helloworld:/${APP_DIR}:/g" docker/docker-compose.tmp.yaml
sed -i "s/APP_DIR: helloworld/APP_DIR: ${APP_DIR}/g" docker/docker-compose.tmp.yaml
sed -i "s/\/app\/helloworld/\/app\/${APP_DIR}/g" docker/docker-compose.tmp.yaml

# Docker環境の起動
cd docker
docker-compose -f docker-compose.tmp.yaml up --build

# 一時ファイルの削除
rm docker-compose.tmp.yaml
