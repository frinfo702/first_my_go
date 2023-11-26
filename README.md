/myproject/                     # プロジェクトのルートディレクトリ
│
├── go.mod                      # Goモジュールの定義ファイル
├── go.sum                      # 依存関係のチェックサムを含むファイル
│
├── main.go                     # メインのGoプログラム
│
├── .git/                       # Gitリポジトリのメタデータ
│
├── .gitignore                  # Gitが無視するファイルのリスト
│
├── README.md                   # プロジェクトの説明ファイル
│
├── pkg/                        # 再利用可能なライブラリコード
│   ├── mylibrary/              
│   │   └── mylibrary.go        
│   │
│   └── anotherlibrary/         
│       └── anotherlibrary.go   
│
└── cmd/                        # プログラムの実行可能ファイル用のソースコード
    └── myapp/                  
        └── main.go             
