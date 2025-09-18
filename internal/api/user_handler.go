package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/quanzhonghuying/GoTaskAI/internal/service"
)

// UserHandler はユーザー関連 API を担当するハンドラ
type UserHandler struct {
    service *service.UserService
}

// NewUserHandler は UserHandler のコンストラクタ
func NewUserHandler(service *service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

// Register ユーザー登録 API
// POST /register
func (h *UserHandler) Register(c *gin.Context) {
    var req struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "不正なリクエスト形式"})
        return
    }

    err := h.service.RegisterUser(req.Name, req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "ユーザー登録成功"})
}

// Login ユーザーログイン API
// POST /login
func (h *UserHandler) Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "不正なリクエスト形式"})
        return
    }

    user, err := h.service.LoginUser(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "ログイン成功",
        "user": gin.H{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
        },
    })
}
