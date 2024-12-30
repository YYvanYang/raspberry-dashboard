package main

import (
	"log"
	"os"

	"github.com/yourusername/raspberry-dashboard/internal/api"
	"github.com/yourusername/raspberry-dashboard/internal/db"
	"github.com/yourusername/raspberry-dashboard/internal/logger"
)

func main() {
	// 确保环境变量已设置
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	// 创建必要的目录
	dirs := []string{"./logs", "./data"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0750); err != nil {
			log.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

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
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	logger.Info("Server starting on :" + port)
	if err := server.Start(":" + port); err != nil {
		logger.Error("Server failed to start: %v", err)
		os.Exit(1)
	}
}
