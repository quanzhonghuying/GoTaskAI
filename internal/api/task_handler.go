package api

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/quanzhonghuying/GoTaskAI/internal/service"
)

// TaskHandler はタスク関連の API を担当する
type TaskHandler struct {
    service *service.TaskService
}

// NewTaskHandler は TaskHandler のコンストラクタ
func NewTaskHandler(service *service.TaskService) *TaskHandler {
    return &TaskHandler{service: service}
}

// CreateTask タスク作成 API
// POST /tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
    var req struct {
        UserID      int    `json:"user_id"`
        Title       string `json:"title"`
        Description string `json:"description"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "不正なリクエスト形式"})
        return
    }

    err := h.service.CreateTask(req.UserID, req.Title, req.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "タスク作成成功"})
}

// ListTasks タスク一覧取得 API
// GET /tasks/:userID
func (h *TaskHandler) ListTasks(c *gin.Context) {
    userIDStr := c.Param("userID")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "無効なユーザーID"})
        return
    }

    tasks, err := h.service.ListTasks(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// MarkTaskCompleted タスク完了 API
// PUT /tasks/:taskID/complete
func (h *TaskHandler) MarkTaskCompleted(c *gin.Context) {
    taskIDStr := c.Param("taskID")
    taskID, err := strconv.Atoi(taskIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "無効なタスクID"})
        return
    }

    err = h.service.MarkTaskCompleted(taskID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "タスク完了済みに更新"})
}
