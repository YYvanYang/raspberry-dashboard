package db

import (
    "database/sql"
    "fmt"
    "os"
    "path/filepath"
    "time"
    _ "github.com/mattn/go-sqlite3"
    "github.com/your-project/utils"
)

type DB struct {
    db *sql.DB
}

// NewDB 创建新的数据库连接
func NewDB(dbPath string) (*DB, error) {
    // 确保数据库目录存在
    dir := filepath.Dir(dbPath)
    if err := os.MkdirAll(dir, 0750); err != nil {
        return nil, fmt.Errorf("failed to create database directory: %w", err)
    }

    // 打开数据库连接
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // 配置连接池
    db.SetMaxOpenConns(1)    // SQLite 限制
    db.SetMaxIdleConns(1)
    db.SetConnMaxLifetime(time.Hour)

    // 启用性能优化
    pragmas := []string{
        "PRAGMA journal_mode=WAL",
        "PRAGMA synchronous=NORMAL",
        "PRAGMA temp_store=MEMORY",
        "PRAGMA mmap_size=30000000000",
        "PRAGMA cache_size=32000",
        "PRAGMA foreign_keys=ON",
    }

    for _, pragma := range pragmas {
        if _, err := db.Exec(pragma); err != nil {
            return nil, fmt.Errorf("failed to set %s: %w", pragma, err)
        }
    }

    // 验证连接
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    return &DB{db: db}, nil
}

// Close 关闭数据库连接
func (db *DB) Close() error {
    return db.db.Close()
}

// InitSchema 初始化数据库架构
func (db *DB) InitSchema() error {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
            password_hash TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`,
        `CREATE TABLE IF NOT EXISTS configs (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            service TEXT UNIQUE NOT NULL,
            config TEXT NOT NULL,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`,
        `CREATE TABLE IF NOT EXISTS logs (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            type TEXT NOT NULL,
            message TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`,
        `CREATE INDEX IF NOT EXISTS idx_logs_created_at ON logs(created_at)`,
        `CREATE INDEX IF NOT EXISTS idx_logs_type ON logs(type)`,
    }

    for _, query := range queries {
        if _, err := db.db.Exec(query); err != nil {
            return fmt.Errorf("failed to execute schema query: %w", err)
        }
    }

    // 检查是否需要创建默认管理员账户
    var count int
    if err := db.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count); err != nil {
        return fmt.Errorf("failed to check users count: %w", err)
    }

    if count == 0 {
        // 创建默认管理员账户
        defaultPassword := "admin123"  // 默认密码
        hashedPassword, err := utils.HashPassword(defaultPassword)
        if err != nil {
            return fmt.Errorf("failed to hash default password: %w", err)
        }

        _, err = db.db.Exec(
            "INSERT INTO users (username, password_hash) VALUES (?, ?)",
            "admin",
            hashedPassword,
        )
        if err != nil {
            return fmt.Errorf("failed to create default admin user: %w", err)
        }
    }

    return nil
}

// Transaction 执行事务
func (db *DB) Transaction(fn func(*sql.Tx) error) error {
    tx, err := db.db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        }
    }()

    if err := fn(tx); err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}

// UpdatePassword 更新用户密码
func (db *DB) UpdatePassword(username string, newPasswordHash string) error {
    result, err := db.db.Exec(
        "UPDATE users SET password_hash = ? WHERE username = ?",
        newPasswordHash,
        username,
    )
    if err != nil {
        return fmt.Errorf("failed to update password: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get affected rows: %w", err)
    }

    if rows == 0 {
        return fmt.Errorf("user not found")
    }

    return nil
}

// VerifyPassword 验证用户当前密码
func (db *DB) VerifyPassword(username string, password string) error {
    var storedHash string
    err := db.db.QueryRow(
        "SELECT password_hash FROM users WHERE username = ?",
        username,
    ).Scan(&storedHash)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("user not found")
        }
        return fmt.Errorf("failed to get user password: %w", err)
    }

    if !utils.CheckPasswordHash(password, storedHash) {
        return fmt.Errorf("invalid password")
    }

    return nil
} 