package db

import (
    "log"

    _ "github.com/lib/pq"     // PostgreSQL ドライバ
    "github.com/jmoiron/sqlx" // sqlx ライブラリ
)

// InitDB: PostgreSQL 接続初期化
func InitDB() (*sqlx.DB, error) {
    dsn := "user=gotaskai_user password=1qaz!QAZ dbname=gotaskai sslmode=disable"

    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("✅ DB 接続成功")
    return db, nil // ✅ 最後に必ず返す
}
