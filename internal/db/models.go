package db

import (
    "time"
    "strings"
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

func (db *DB) GetUserByID(id int64) (*User, error) {
    var user User
    err := db.db.QueryRow("SELECT id, username, password_hash, created_at FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (db *DB) GetUserByUsername(username string) (*User, error) {
    var user User
    err := db.db.QueryRow("SELECT id, username, password_hash, created_at FROM users WHERE username = ?", username).
        Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (db *DB) UpdateUserPassword(userID int64, passwordHash string) error {
    _, err := db.db.Exec("UPDATE users SET password_hash = ? WHERE id = ?", passwordHash, userID)
    return err
}

func (db *DB) GetAllConfigs() (map[string]string, error) {
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

func (db *DB) UpdateConfig(service, config string) error {
    _, err := db.db.Exec(
        `INSERT INTO configs (service, config) VALUES (?, ?)
         ON CONFLICT(service) DO UPDATE SET config = ?, updated_at = CURRENT_TIMESTAMP`,
        service, config, config,
    )
    return err
}

func (db *DB) GetFilteredLogs(logType, search string, since time.Time) ([]Log, error) {
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

func (db *DB) AddLog(logType, message string) error {
    _, err := db.db.Exec(
        "INSERT INTO logs (type, message) VALUES (?, ?)",
        strings.ToLower(logType),
        message,
    )
    return err
} 