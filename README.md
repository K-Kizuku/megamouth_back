# ハックツハッカソン メガマウスカップ バックエンド

## Getting Start!
### 起動方法
Dockerを起動して  
`make up`  
でAPIサーバが起動します  
`make down`  
でAPIサーバが止まります  
http://localhost:8000 で動作しています

### 開発方法
サーバ起動後  
http://localhost:8000/swagger/index.html → OpenAPI  
http://localhost:8080 → adminer  
が動作しています  

## 実装詳細
バックエンドの実装の詳細について
### DB定義
下記URL参照  
https://dbdiagram.io/d/64b0424302bd1c4a5e0a3a7d

### ディレクトリ構成について
(オレオレ)クリーンアーキテクチャを採用しています
```
.
├── Dockerfile
├── Makefile
├── README.md
├── api
│   ├── adapter // interfaceのメソッドの提供
│   │   ├── controller // 入力に対するアダプター
│   │   │   └── user.go
│   │   ├── gateway // DB操作に対するアダプター
│   │   │   └── user.go
│   │   └── presenter // 出力に対するアダプター
│   │       └── user.go
│   ├── driver // 技術的な実装を持つ
│   │   ├── db // DBの接続
│   │   │   └── db.go
│   │   ├── server.go // サーバの起動
│   │   └── user.go
│   ├── entity // ドメインモデルを実装
│   │   ├── models // DBのモデル定義
│   │   │   └── user.go
│   │   └── repository // CRUDs
│   │       └── user.go
│   ├── usecase // ビジネスロジックを実行
│   │   ├── interactor // inputとoutputの接続
│   │   │   └── user.go
│   │   ├── port // 入出力のポートの提供
│   │   │   └── user.go
│   │   └── schema // スキーマ定義
│   │       └── user.go
│   └── utils // その他の汎用性の高い関数群
│       ├── codes // PJ内で使うcode群
│       │   └── codes.go
│       ├── config // 環境変数の読み込み
│       │   └── env.go
│       ├── errors // 独自のエラー処理
│       │   └── errors.go
│       ├── middleware // ミドルウェア
│       │   └── authMiddleware.go
│       └── tools // 各種ツール群
│           ├── hash.go // ハッシュ化
│           └── token.go // JWT発行
├── compose.yml
├── docs // OpenAPI用．swaggo/swagによる自動生成
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── grpc // gRPCの自動生成ファイル
│   ├── recognizer.pb.go
│   ├── recognizer.proto
│   └── recognizer_grpc.pb.go
├── go.mod
├── go.sum
└──  main.go
```
