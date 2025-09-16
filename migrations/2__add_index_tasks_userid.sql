-- +goose Up
-- 理由: ユーザーごとのタスク一覧取得を高速化するため user_id にインデックスを追加

CREATE INDEX idx_tasks_user_id ON tasks (user_id);

-- +goose Down
-- 理由: ロールバック時にインデックスを削除
DROP INDEX IF EXISTS idx_tasks_user_id;