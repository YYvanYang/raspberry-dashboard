package utils

import (
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// PasswordPolicy 定义密码策略
type PasswordPolicy struct {
	MinLength      int
	RequireUpper   bool
	RequireLower   bool
	RequireNumber  bool
	RequireSpecial bool
}

// DefaultPasswordPolicy 默认密码策略
var DefaultPasswordPolicy = PasswordPolicy{
	MinLength:      8,
	RequireUpper:   true,
	RequireLower:   true,
	RequireNumber:  true,
	RequireSpecial: true,
}

// ValidatePassword 验证密码是否符合策略
func ValidatePassword(password string, policy PasswordPolicy) error {
	if len(password) < policy.MinLength {
		return fmt.Errorf("密码长度至少需要 %d 个字符", policy.MinLength)
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if policy.RequireUpper && !hasUpper {
		return fmt.Errorf("密码必须包含大写字母")
	}
	if policy.RequireLower && !hasLower {
		return fmt.Errorf("密码必须包含小写字母")
	}
	if policy.RequireNumber && !hasNumber {
		return fmt.Errorf("密码必须包含数字")
	}
	if policy.RequireSpecial && !hasSpecial {
		return fmt.Errorf("密码必须包含特殊字符")
	}

	return nil
}

// HashPassword 生成密码哈希
func HashPassword(password string) (string, error) {
	// 先验证密码策略
	if err := ValidatePassword(password, DefaultPasswordPolicy); err != nil {
		return "", err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

// CheckPasswordHash 验证密码哈希
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
