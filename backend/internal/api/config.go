package api

import (
    "encoding/json"
    "github.com/gofiber/fiber/v2"
)

// 配置处理函数
func (s *Server) handleGetConfigs(c *fiber.Ctx) error {
    configs, err := s.db.GetAllConfigs()
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to get configs")
    }
    return c.JSON(configs)
}

func (s *Server) handleGetConfig(c *fiber.Ctx) error {
    service := c.Params("service")
    config, err := s.db.GetConfig(service)
    if err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Config not found")
    }
    return c.JSON(config)
}

func (s *Server) handleUpdateConfig(c *fiber.Ctx) error {
    service := c.Params("service")
    
    var configData json.RawMessage
    if err := c.BodyParser(&configData); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid config data")
    }

    if err := s.db.SaveConfig(service, string(configData)); err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to save config")
    }

    // 记录配置更新日志
    s.db.AddLog("config", "Updated configuration for "+service)
    
    // 尝试重启相关服务
    if err := s.service.RestartService(service); err != nil {
        s.db.AddLog("warning", "Failed to restart service "+service+": "+err.Error())
    }

    return c.SendStatus(fiber.StatusOK)
} 