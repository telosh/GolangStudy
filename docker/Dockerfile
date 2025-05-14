# ベースイメージとして公式のGolangイメージを使用
FROM golang:1.20-alpine

# 作業ディレクトリを設定
WORKDIR /app

# Go Modulesの使用を許可
ENV GO111MODULE=on

# ソースコードをコンテナにコピー
COPY . .

# アプリケーションをビルド
# ビルド対象のディレクトリは環境変数で指定可能
ARG APP_DIR=helloworld
RUN cd ${APP_DIR} && go build -o ../app main.go

# 実行可能ファイルをデフォルトのコマンドとして設定
CMD ["./app"]