package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/quanzhonghuying/GoTaskAI/internal/model"
)

// TaskRepository は tasks テーブルへの操作を担当する
type TaskRepository struct {
    db *sqlx.DB
}

// コンストラクタ
func NewTaskRepository(db *sqlx.DB) *TaskRepository {
    return &TaskRepository{db: db}
}

// Create 新しいタスクをDBに保存
func (r *TaskRepository) Create(task *model.Task) error {
    _, err := r.db.Exec(
        "INSERT INTO tasks (user_id, title, description, completed) VALUES ($1, $2, $3, $4)",
        task.UserID, task.Title, task.Description, task.Completed,
    )
    return err
}

// MarkCompleted タスクを完了状態に更新
func (r *TaskRepository) MarkCompleted(taskID int) error {
    _, err := r.db.Exec(
        "UPDATE tasks SET completed=true WHERE id=$1",
        taskID,
    )
    return err
}

// FindByUserID ユーザーIDに紐づくタスク一覧を取得
func (r *TaskRepository) FindByUserID(userID int) ([]model.Task, error) {
    var tasks []model.Task
    err := r.db.Select(&tasks, "SELECT * FROM tasks WHERE user_id=$1", userID)
    return tasks, err
}
