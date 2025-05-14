JSONファイルの書き込み:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"operation":"write","filename":"data/people.json","format":"json","data":[{"name":"John Doe","age":30,"email":"john@example.com"}]}' http://localhost:8083/file
```

CSVファイルの書き込み:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"operation":"write","filename":"data/people.csv","format":"csv","data":[{"name":"John Doe","age":30,"email":"john@example.com"}]}' http://localhost:8083/file
```

ファイルの読み込み:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"operation":"read","filename":"data/people.json","format":"json"}' http://localhost:8083/file
```

ファイルへの追記:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"operation":"append","filename":"data/people.json","format":"json","data":{"name":"Jane Smith","age":25,"email":"jane@example.com"}}' http://localhost:8083/file
```

ディレクトリ内のファイル一覧取得:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"operation":"list","filename":"data"}' http://localhost:8083/file
```

テキストファイルの書き込み:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"operation":"write","filename":"data/note.txt","format":"text","data":"Hello, World!"}' http://localhost:8083/file
```

テキストファイルへの追記:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"operation":"append","filename":"data/note.txt","format":"text","data":"\nThis is a new line."}' http://localhost:8083/file
```

こ