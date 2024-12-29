package api

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/yourusername/raspberry-dashboard/internal/db"
    "github.com/yourusername/raspberry-dashboard/internal/service"
)

type Server struct {
    app     *fiber.App
    db      *db.Database
    service *service.SystemService
}

func NewServer(database *db.Database) *Server {
    app := fiber.New(fiber.Config{
        ErrorHandler: errorHandler,
    })

    // 中间件
    app.Use(logger.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowMethods: "GET, POST, PUT, DELETE",
    }))

    server := &Server{
        app:     app,
        db:      database,
        service: service.NewSystemService(),
    }

    // 设置路由
    server.setupRoutes()

    return server
}

func (s *Server) Start(addr string) error {
    return s.app.Listen(addr)
}

func (s *Server) setupRoutes() {
    // 公开路由
    s.app.Post("/api/auth/login", s.handleLogin)

    // 需要认证的路由
    api := s.app.Group("/api", s.authMiddleware)
    
    // 认证相关
    auth := api.Group("/auth")
    auth.Post("/logout", s.handleLogout)
    auth.Get("/profile", s.handleGetProfile)
    
    // 配置管理
    configs := api.Group("/configs")
    configs.Get("/", s.handleGetConfigs)
    configs.Get("/:service", s.handleGetConfig)
    configs.Put("/:service", s.handleUpdateConfig)
    
    // 系统管理
    system := api.Group("/system")
    system.Get("/status", s.handleGetStatus)
    system.Get("/logs", s.handleGetLogs)
    system.Post("/services/:name", s.handleControlService)
} 