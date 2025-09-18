package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/quanzhonghuying/GoTaskAI/internal/db"
    "github.com/quanzhonghuying/GoTaskAI/internal/repository"
    "github.com/quanzhonghuying/GoTaskAI/internal/service"
    "github.com/quanzhonghuying/GoTaskAI/internal/api"
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
    r.POST("/register", userHandler.Register)
    r.POST("/login", userHandler.Login)
    r.POST("/tasks", taskHandler.CreateTask)
    r.GET("/tasks/:userID", taskHandler.ListTasks)
    r.PUT("/tasks/:taskID/complete", taskHandler.MarkTaskCompleted)

    // サーバー起動
    r.Run(":8080")
}
