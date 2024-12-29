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

// changePasswordRequest 修改密码请求
type changePasswordRequest struct {
    CurrentPassword string `json:"currentPassword"`
    NewPassword     string `json:"newPassword"`
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

// handleChangePassword 处理修改密码请求
func (s *Server) handleChangePassword(c *fiber.Ctx) error {
    // 获取当前用户信息
    claims := c.Locals("user").(*utils.Claims)
    
    // 解析请求
    var req changePasswordRequest
    if err := c.BodyParser(&req); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
    }

    // 验证密码格式
    if len(req.NewPassword) < 8 {
        return fiber.NewError(fiber.StatusBadRequest, "Password must be at least 8 characters long")
    }

    // 验证当前密码
    if err := s.db.VerifyPassword(claims.Username, req.CurrentPassword); err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Current password is incorrect")
    }

    // 生成新密码哈希
    newPasswordHash, err := utils.HashPassword(req.NewPassword)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash new password")
    }

    // 更新密码
    if err := s.db.UpdatePassword(claims.Username, newPasswordHash); err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to update password")
    }

    return c.SendStatus(fiber.StatusOK)
} 