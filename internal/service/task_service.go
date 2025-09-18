package service

import (
    "github.com/quanzhonghuying/GoTaskAI/internal/model"
    "github.com/quanzhonghuying/GoTaskAI/internal/repository"
)

// TaskService はタスクに関するビジネスロジックを担当する
type TaskService struct {
    repo *repository.TaskRepository
}

// NewTaskService は TaskService のコンストラクタ
func NewTaskService(repo *repository.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

// CreateTask 新しいタスクを作成する
func (s *TaskService) CreateTask(userID int, title, description string) error {
    task := &model.Task{
        UserID:      userID,
        Title:       title,
        Description: description,
        Completed:   false,
    }
    return s.repo.Create(task)
}

// MarkTaskCompleted タスクを完了状態にする
func (s *TaskService) MarkTaskCompleted(taskID int) error {
    return s.repo.MarkCompleted(taskID)
}

// ListTasks 指定ユーザーのタスク一覧を取得する
func (s *TaskService) ListTasks(userID int) ([]model.Task, error) {
    return s.repo.FindByUserID(userID)
}
