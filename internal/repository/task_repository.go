package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/quanzhonghuying/GoTaskAI/internal/model"
)

// TaskRepository は tasks テーブルへのアクセスを担当する
type TaskRepository struct {
    db *sqlx.DB
}

// タスク新規作成
func (r *TaskRepository) CreateTask(userID int, title, description string) error {
    _, err := r.db.Exec(
        "INSERT INTO tasks (user_id, title, description, completed) VALUES ($1, $2, $3, $4)",
        userID, title, description, false,
    )
    return err
}

// ユーザーのタスク一覧を取得
func (r *TaskRepository) FindTasksByUser(userID int) ([]model.Task, error) {
    var tasks []model.Task
    err := r.db.Select(&tasks, "SELECT * FROM tasks WHERE user_id=$1", userID)
    if err != nil {
        return nil, err
    }
    return tasks, nil
}

// タスクを完了済みに更新
func (r *TaskRepository) MarkTaskCompleted(taskID int) error {
    _, err := r.db.Exec("UPDATE tasks SET completed=true WHERE id=$1", taskID)
    return err
}
