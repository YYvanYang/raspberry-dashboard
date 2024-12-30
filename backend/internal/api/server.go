package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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
		ErrorHandler:            errorHandler,
		EnableTrustedProxyCheck: false,
	})

	// 中间件
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
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
	protected := api.Group("", s.authMiddleware)

	// 认证相关
	auth := protected.Group("/auth")
	auth.Post("/logout", s.handleLogout)
	auth.Get("/profile", s.handleGetProfile)
	auth.Post("/change-password", s.handleChangePassword)

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
		Root:         http.Dir("web"),
		Browse:       false,
		Index:        "index.html",
		MaxAge:       86400,
		NotFoundFile: "index.html",
	}))
}
