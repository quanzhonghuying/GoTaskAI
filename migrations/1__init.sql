-- +goose Up
-- 初期スキーマ作成: users / tasks テーブル

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, -- ユーザーID
    email VARCHAR(255) UNIQUE NOT NULL, -- メールアドレス（一意制約）
    password_hash VARCHAR(255) NOT NULL, -- ハッシュ化パスワード
    created_at TIMESTAMP DEFAULT now() -- 作成日時
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY, -- タスクID
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE, -- 所属ユーザー
    title VARCHAR(255) NOT NULL, -- タスクタイトル
    description TEXT, -- タスク詳細
    completed BOOLEAN DEFAULT false, -- 完了フラグ
    created_at TIMESTAMP DEFAULT now(), -- 作成日時
    updated_at TIMESTAMP DEFAULT now() -- 更新日時
);

-- +goose Down
-- テーブル削除（ロールバック用）

DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS users;