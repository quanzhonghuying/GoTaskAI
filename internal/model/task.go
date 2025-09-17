package model

import "time"

// Task 構造体は tasks テーブルの1行を表す
type Task struct {
    ID          int       `db:"id" json:"id"`                   // 主キーID
    UserID      int       `db:"user_id" json:"user_id"`         // 紐づくユーザーID（外部キー）
    Title       string    `db:"title" json:"title"`             // タスクタイトル
    Description string    `db:"description" json:"description"` // タスク詳細
    Completed   bool      `db:"completed" json:"completed"`     // 完了フラグ
    CreatedAt   time.Time `db:"created_at" json:"created_at"`   // 作成日時
    UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`   // 更新日時
}
