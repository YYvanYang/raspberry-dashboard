package api

import (
    "strings"
    "github.com/gofiber/fiber/v2"
    "github.com/yourusername/raspberry-dashboard/internal/utils"
)

func (s *Server) authMiddleware(c *fiber.Ctx) error {
    // 获取 Authorization header
    auth := c.Get("Authorization")
    if auth == "" {
        return fiber.NewError(fiber.StatusUnauthorized, "Missing authorization header")
    }

    // 检查 Bearer token
    parts := strings.Split(auth, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return fiber.NewError(fiber.StatusUnauthorized, "Invalid authorization header")
    }

    // 验证 token
    token := parts[1]
    claims, err := utils.ValidateToken(token)
    if err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
    }

    // 将用户信息存储在上下文中
    c.Locals("user", claims)
    return c.Next()
} 