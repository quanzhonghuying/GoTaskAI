package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/your_project/internal/model"
)

// UserRepository は users テーブルへのアクセスを担当する
type UserRepository struct {
    db *sqlx.DB
}

// メールアドレスでユーザーを検索
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
    var user model.User
    err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// 新規ユーザー作成
func (r *UserRepository) Create(user *model.User) error {
    _, err := r.db.Exec(
        "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
        user.Name, user.Email, user.Password,
    )
    return err
}

// IDでユーザーを検索
func (r *UserRepository) FindByID(id int) (*model.User, error) {
    var user model.User
    err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
