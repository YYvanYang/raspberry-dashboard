package main

import (
    "log"
    "os"
    
    "github.com/yourusername/raspberry-dashboard/internal/api"
    "github.com/yourusername/raspberry-dashboard/internal/db"
    "github.com/yourusername/raspberry-dashboard/internal/logger"
)

func main() {
    // 初始化日志
    if err := logger.Init("./logs/app.log"); err != nil {
        log.Fatalf("Failed to initialize logger: %v", err)
    }

    // 初始化数据库
    database, err := db.NewDB("./data/dashboard.db")
    if err != nil {
        logger.Error("Failed to initialize database: %v", err)
        os.Exit(1)
    }
    defer database.Close()

    // 初始化数据库架构
    if err := database.InitSchema(); err != nil {
        logger.Error("Failed to initialize database schema: %v", err)
        os.Exit(1)
    }

    // 记录启动日志
    logger.Info("Application started successfully")

    // 启动服务器
    server := api.NewServer(database)
    logger.Info("Server starting on :3001")
    if err := server.Start(":3001"); err != nil {
        logger.Error("Server failed to start: %v", err)
        os.Exit(1)
    }
} 