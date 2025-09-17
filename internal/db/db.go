package db

import (
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

// DB 接続オブジェクト（アプリ全体で共有）
var DB *sqlx.DB

// InitDB: PostgreSQL 接続初期化
func InitDB() {
    dsn := "user=gotaskai_user password=1qaz!QAZ dbname=gotaskai sslmode=disable"

    db, err := sqlx.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("DB 接続失敗: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("DB 通信確認失敗: %v", err)
    }

    DB = db
}
