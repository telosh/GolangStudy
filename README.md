# Golang Study Project

このプロジェクトは、Golangの学習用のDocker環境を提供します。

## 前提条件

- Docker
- Docker Compose
- Git

## 使用方法

1. プロジェクトをクローンします：
```bash
git clone <repository-url>
cd GolangStudy
```

2. アプリケーションを起動します：
```bash
./start.sh <app_directory>
```

例：
```bash
./start.sh helloworld
```

## プロジェクト構造

```
.
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yaml
├── helloworld/
│   └── main.go
├── practice-api/
│   └── ...
├── start.sh
└── README.md
```

## 注意事項

- `start.sh`を実行する前に、実行権限を付与してください：
```bash
chmod +x start.sh
```

- アプリケーションディレクトリは、`main.go`を含む必要があります。

## 開発

新しいアプリケーションを追加する場合は、以下の手順に従ってください：

1. 新しいディレクトリを作成し、`main.go`を配置します。
2. `start.sh`を使用して、新しいアプリケーションを起動します。

例：
```bash
mkdir myapp
# main.goを作成
./start.sh myapp
```
