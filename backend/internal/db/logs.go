package db

import (
    "time"
    "strings"
)

type Log struct {
    ID        int64     `json:"id"`
    Type      string    `json:"type"`
    Message   string    `json:"message"`
    CreatedAt time.Time `json:"created_at"`
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