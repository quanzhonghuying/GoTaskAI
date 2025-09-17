package main

import (
    "fmt"
    "github.com/quanzhonghuying/GoTaskAI/internal/db"
)

func main() {
    // データベース接続を初期化
    db.InitDB()

    // DB オブジェクトが有効か確認
    if db.DB != nil {
        fmt.Println("✅ DB オブジェクト準備完了")
    }
}
