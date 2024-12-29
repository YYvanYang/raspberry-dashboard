package api

import (
    "github.com/gofiber/fiber/v2"
    "github.com/yourusername/raspberry-dashboard/internal/monitor"
)

func (s *Server) handleGetMetrics(c *fiber.Ctx) error {
    metrics, err := monitor.GetSystemMetrics()
    if err != nil {
        return &AppError{
            Code:    fiber.StatusInternalServerError,
            Message: "Failed to get system metrics",
            Err:     err,
        }
    }

    return c.JSON(metrics)
}

func (s *Server) handleHealthCheck(c *fiber.Ctx) error {
    // 检查关键服务状态
    services := []string{"hysteria2", "wg-quick@wg0", "cloudflared"}
    serviceStatus := make(map[string]bool)

    for _, service := range services {
        active, _ := monitor.CheckServiceStatus(service)
        serviceStatus[service] = active
    }

    // 检查系统资源
    metrics, err := monitor.GetSystemMetrics()
    if err != nil {
        return &AppError{
            Code:    fiber.StatusInternalServerError,
            Message: "Failed to check system health",
            Err:     err,
        }
    }

    // 定义健康阈值
    healthStatus := "healthy"
    if metrics.CPU > 90 || metrics.Memory.UsageRate > 90 || metrics.Disk.UsageRate > 90 {
        healthStatus = "warning"
    }

    return c.JSON(fiber.Map{
        "status":   healthStatus,
        "services": serviceStatus,
        "metrics":  metrics,
    })
} 