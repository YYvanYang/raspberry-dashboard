package api

import (
    "github.com/gofiber/fiber/v2"
    "github.com/yourusername/raspberry-dashboard/internal/service"
)

type SystemStatus struct {
    CPU       float64            `json:"cpu"`
    Memory    service.MemoryInfo `json:"memory"`
    Disk      service.DiskInfo   `json:"disk"`
    Uptime    string            `json:"uptime"`
    Services  []ServiceStatus    `json:"services"`
}

type ServiceStatus struct {
    Name    string `json:"name"`
    Status  string `json:"status"`
    Running bool   `json:"running"`
}

func (s *Server) handleGetStatus(c *fiber.Ctx) error {
    status, err := s.service.GetSystemStatus()
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to get system status")
    }
    return c.JSON(status)
}

func (s *Server) handleGetLogs(c *fiber.Ctx) error {
    limit := 100 // 默认获取最近100条日志
    logs, err := s.db.GetLogs(limit)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to get logs")
    }
    return c.JSON(logs)
}

func (s *Server) handleControlService(c *fiber.Ctx) error {
    serviceName := c.Params("name")
    action := c.Query("action", "restart") // 默认操作为重启

    switch action {
    case "start":
        err := s.service.StartService(serviceName)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to start service")
        }
    case "stop":
        err := s.service.StopService(serviceName)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to stop service")
        }
    case "restart":
        err := s.service.RestartService(serviceName)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to restart service")
        }
    default:
        return fiber.NewError(fiber.StatusBadRequest, "Invalid action")
    }

    return c.SendStatus(fiber.StatusOK)
} 