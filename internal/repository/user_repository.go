package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/quanzhonghuying/GoTaskAI/internal/model"
)

// UserRepository は users テーブルへのアクセスを担当する
type UserRepository struct {
    db *sqlx.DB
}

// コンストラクタ
func NewUserRepository(db *sqlx.DB) *UserRepository {
    return &UserRepository{db: db}
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
// Create ユーザー新規登録
func (r *UserRepository) Create(user *model.User) error {
    _, err := r.db.Exec(
        "INSERT INTO users (email, password_hash) VALUES ($1, $2)",
        user.Email, user.PasswordHash,
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
