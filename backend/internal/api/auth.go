package api

import (
    "github.com/gofiber/fiber/v2"
    "github.com/yourusername/raspberry-dashboard/internal/utils"
)

type loginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type loginResponse struct {
    Token string `json:"token"`
}

func (s *Server) handleLogin(c *fiber.Ctx) error {
    var req loginRequest
    if err := c.BodyParser(&req); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
    }

    // 获取用户
    user, err := s.db.GetUserByUsername(req.Username)
    if err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
    }

    // 验证密码
    if !utils.CheckPasswordHash(req.Password, user.Password) {
        return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
    }

    // 生成 token
    token, err := utils.GenerateToken(user.ID, user.Username)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token")
    }

    return c.JSON(loginResponse{Token: token})
}

func (s *Server) handleLogout(c *fiber.Ctx) error {
    // 实际应用中可能需要将token加入黑名单
    return c.SendStatus(fiber.StatusOK)
}

func (s *Server) handleGetProfile(c *fiber.Ctx) error {
    claims := c.Locals("user").(*utils.Claims)
    user, err := s.db.GetUserByUsername(claims.Username)
    if err != nil {
        return fiber.NewError(fiber.StatusNotFound, "User not found")
    }

    return c.JSON(user)
} 