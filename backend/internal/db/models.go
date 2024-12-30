package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// 用户模型
type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

// 配置模型
type Config struct {
	ID        int64     `json:"id"`
	Service   string    `json:"service"`
	Config    string    `json:"config"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 日志模型
type Log struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

// User 相关方法
func (db *Database) GetUserByID(id int64) (*User, error) {
	var user User
	err := db.db.QueryRow("SELECT id, username, password_hash, created_at FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *Database) GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.db.QueryRow("SELECT id, username, password_hash, created_at FROM users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *Database) UpdateUserPassword(userID int64, passwordHash string) error {
	_, err := db.db.Exec("UPDATE users SET password_hash = ? WHERE id = ?", passwordHash, userID)
	return err
}

// Config 相关方法
func (db *Database) GetAllConfigs() (map[string]string, error) {
	rows, err := db.db.Query("SELECT service, config FROM configs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	configs := make(map[string]string)
	for rows.Next() {
		var service, config string
		if err := rows.Scan(&service, &config); err != nil {
			return nil, err
		}
		configs[service] = config
	}
	return configs, nil
}

func (db *Database) UpdateConfig(service, config string) error {
	_, err := db.db.Exec(
		`INSERT INTO configs (service, config) VALUES (?, ?)
		 ON CONFLICT(service) DO UPDATE SET config = ?, updated_at = CURRENT_TIMESTAMP`,
		service, config, config,
	)
	return err
}

// Log 相关方法
func (db *Database) GetFilteredLogs(logType, search string, since time.Time) ([]Log, error) {
	query := `SELECT id, type, message, created_at FROM logs WHERE created_at >= ?`
	args := []interface{}{since}

	if logType != "" && logType != "all" {
		query += ` AND type = ?`
		args = append(args, logType)
	}

	if search != "" {
		query += ` AND message LIKE ?`
		args = append(args, "%"+search+"%")
	}

	query += ` ORDER BY created_at DESC LIMIT 1000`

	rows, err := db.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []Log
	for rows.Next() {
		var log Log
		err := rows.Scan(&log.ID, &log.Type, &log.Message, &log.CreatedAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil
}

func (db *Database) AddLog(logType, message string) error {
	_, err := db.db.Exec(
		"INSERT INTO logs (type, message) VALUES (?, ?)",
		strings.ToLower(logType),
		message,
	)
	return err
}

// GetConfig 获取指定服务的配置
func (db *Database) GetConfig(service string) (*Config, error) {
	var config Config
	err := db.db.QueryRow(
		"SELECT id, service, config, updated_at FROM configs WHERE service = ?",
		service,
	).Scan(&config.ID, &config.Service, &config.Config, &config.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("config not found for service: %s", service)
		}
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	return &config, nil
}

// SaveConfig 保存服务配置
func (db *Database) SaveConfig(service, configData string) error {
	_, err := db.db.Exec(
		`INSERT INTO configs (service, config) VALUES (?, ?)
		ON CONFLICT(service) DO UPDATE SET config = ?, updated_at = CURRENT_TIMESTAMP`,
		service, configData, configData,
	)
	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}
	return nil
}

// GetLogs 获取指定数量的最新日志
func (db *Database) GetLogs(limit int) ([]Log, error) {
	rows, err := db.db.Query(
		"SELECT id, type, message, created_at FROM logs ORDER BY created_at DESC LIMIT ?",
		limit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}
	defer rows.Close()

	var logs []Log
	for rows.Next() {
		var log Log
		err := rows.Scan(&log.ID, &log.Type, &log.Message, &log.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan log: %w", err)
		}
		logs = append(logs, log)
	}

	return logs, nil
}
