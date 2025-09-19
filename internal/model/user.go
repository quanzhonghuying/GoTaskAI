package model

// User 構造体は users テーブルの1行を表す
type User struct {
    ID           int    `db:"id"`
    Email        string `db:"email"`
    PasswordHash string `db:"password_hash"`
    CreatedAt    string `db:"created_at"`
}

