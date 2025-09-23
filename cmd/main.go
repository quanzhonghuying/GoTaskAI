package main

import (
    "log"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/quanzhonghuying/GoTaskAI/internal/api"
    "github.com/quanzhonghuying/GoTaskAI/internal/db"
    "github.com/quanzhonghuying/GoTaskAI/internal/repository"
    "github.com/quanzhonghuying/GoTaskAI/internal/service"
)

func main() {
    // DB 初期化
    database, err := db.InitDB()
    if err != nil {
        log.Fatal("DB 初期化失敗: ", err)
    }

    // Repository 初期化
    userRepo := repository.NewUserRepository(database)
    taskRepo := repository.NewTaskRepository(database)

    // Service 初期化
    userService := service.NewUserService(userRepo)
    taskService := service.NewTaskService(taskRepo)

    // Handler 初期化
    userHandler := api.NewUserHandler(userService)
    taskHandler := api.NewTaskHandler(taskService)

    // Gin ルーター設定
    r := gin.Default()

    // ✅ CORS 設定追加
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // ルート設定
    r.POST("/register", userHandler.Register)
    r.POST("/login", userHandler.Login)
    r.POST("/tasks", taskHandler.CreateTask)
    r.GET("/tasks/:userID", taskHandler.ListTasks)
    r.PUT("/tasks/:taskID/complete", taskHandler.MarkTaskCompleted)

    // サーバー起動
    r.Run(":8080")
}
