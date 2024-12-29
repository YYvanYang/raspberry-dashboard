package api

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// AppError 定义应用错误类型
type AppError struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Details interface{} `json:"details,omitempty"`
    Err     error      `json:"-"` // 内部错误不暴露给客户端
}

func (e *AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

// NewError 创建新的应用错误
func NewError(code int, message string, err error) *AppError {
    return &AppError{
        Code:    code,
        Message: message,
        Err:     err,
    }
}

// ErrorHandler 全局错误处理中间件
func ErrorHandler(c *fiber.Ctx, err error) error {
    // 默认错误响应
    code := fiber.StatusInternalServerError
    message := "Internal Server Error"
    var details interface{}

    // 处理应用错误
    if e, ok := err.(*AppError); ok {
        code = e.Code
        message = e.Message
        details = e.Details
    } else if e, ok := err.(*fiber.Error); ok {
        // 处理 Fiber 内置错误
        code = e.Code
        message = e.Message
    }

    // 开发环境下记录详细错误
    if fiber.IsChild() {
        fmt.Printf("Error: %+v\n", err)
    }

    // 返回 JSON 格式错误响应
    return c.Status(code).JSON(fiber.Map{
        "error": message,
        "details": details,
    })
} 