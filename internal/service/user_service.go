package service

import (
    "errors"

    "golang.org/x/crypto/bcrypt"
    "github.com/quanzhonghuying/GoTaskAI/internal/model"
    "github.com/quanzhonghuying/GoTaskAI/internal/repository"
)

// UserService はユーザーに関するビジネスロジックを担当する
type UserService struct {
    repo *repository.UserRepository
}

// NewUserService は UserService のコンストラクタ
func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

// RegisterUser ユーザー登録処理
// - パスワードを bcrypt でハッシュ化
// - DB に保存
func (s *UserService) RegisterUser(name, email, password string) error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &model.User{
        Name:     name,
        Email:    email,
        Password: string(hashed),
    }

    return s.repo.Create(user)
}

// LoginUser ユーザーログイン処理
// - email でユーザー検索
// - 入力パスワードとハッシュを照合
func (s *UserService) LoginUser(email, password string) (*model.User, error) {
    user, err := s.repo.FindByEmail(email)
    if err != nil {
        return nil, errors.New("ユーザーが存在しません")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, errors.New("パスワードが正しくありません")
    }

    return user, nil
}
