package model

// User 構造体は users テーブルの1行を表す
type User struct {
    ID       int    `db:"id" json:"id"`       // 主キーID
    Name     string `db:"name" json:"name"`   // ユーザー名
    Email    string `db:"email" json:"email"` // メールアドレス
    Password string `db:"password" json:"-"`  // パスワード（JSON出力時は除外）
}
