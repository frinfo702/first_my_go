/myproject/                     # プロジェクトのルートディレクトリ
│
├── go.mod                      # Goモジュールの定義ファイル
├── go.sum                      # 依存関係のチェックサムを含むファイル
│
├── main.go                     # メインのGoプログラム
│
├── .git/                       # Gitリポジトリのメタデータ
├── .gitignore                  # Gitが無視するファイルのリスト
│
├── README.md                   # プロジェクトの説明ファイル
│
├── docs/                       # 学習ノートやドキュメント
│   ├── learning_note.md
│   └── another_note.md
│
├── experiments/                # 実験的なコードや練習用スクリプト
│   ├── experiment1.go
│   └── experiment2.go
│
├── pkg/                        # 再利用可能なライブラリコード
│   ├── mylibrary/              
│   │   └── mylibrary.go        
│   ├── anotherlibrary/         
│   │   └── anotherlibrary.go   
│   └── samples/                # 特定の機能やライブラリのサンプルコード
│       └── sample.go
│
└── cmd/                        # プログラムの実行可能ファイル用のソースコード
    ├── myapp/                  
    │   └── main.go             
    └── smallapp/               # 練習用の小規模なアプリケーション
        └── main.go
