package api

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/filesystem"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/yourusername/raspberry-dashboard/internal/db"
    "github.com/yourusername/raspberry-dashboard/internal/service"
    "net/http"
)

type Server struct {
    app     *fiber.App
    db      *db.Database
    service *service.SystemService
}

func NewServer(database *db.Database) *Server {
    app := fiber.New(fiber.Config{
        ErrorHandler: errorHandler,
        EnableTrustedProxyCheck: false,
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
    // API 路由组
    api := s.app.Group("/api")

    // 公开路由
    api.Post("/auth/login", s.handleLogin)

    // 需要认证的路由
    protected := api.Group("/", s.authMiddleware)
    
    // 认证相关
    auth := protected.Group("/auth")
    auth.Post("/logout", s.handleLogout)
    auth.Get("/profile", s.handleGetProfile)
    
    // 配置管理
    configs := protected.Group("/configs")
    configs.Get("/", s.handleGetConfigs)
    configs.Get("/:service", s.handleGetConfig)
    configs.Put("/:service", s.handleUpdateConfig)
    
    // 系统管理
    system := protected.Group("/system")
    system.Get("/status", s.handleGetStatus)
    system.Get("/logs", s.handleGetLogs)
    system.Post("/services/:name", s.handleControlService)

    // 静态文件服务
    s.app.Use("/", filesystem.New(filesystem.Config{
        Root:         http.Dir("web"),          // 前端构建文件目录
        Browse:       false,                    // 禁止目录浏览
        Index:        "index.html",             // 默认文件
        MaxAge:       86400,                    // 缓存时间：1天
        NotFoundFile: "index.html",             // SPA 支持：所有未匹配路由返回 index.html
    }))
} 