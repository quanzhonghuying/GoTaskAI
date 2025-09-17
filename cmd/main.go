package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL ドライバ
)

func main() {
	// DSN (Data Source Name)
	dsn := "user=gotaskai_user password=1qaz!QAZ dbname=gotaskai sslmode=disable"

	// 1. データベースに接続
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}
	defer db.Close()

	fmt.Println("✅ DB接続成功")

	// 2. 簡単なクエリを実行して確認
	var currentTime string
	err = db.Get(&currentTime, "SELECT NOW()")
	if err != nil {
		log.Fatalf("クエリ実行失敗: %v", err)
	}

	fmt.Println("現在時刻:", currentTime)
}
