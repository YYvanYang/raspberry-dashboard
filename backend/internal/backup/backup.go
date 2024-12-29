package backup

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "sort"
    "time"
)

type BackupConfig struct {
    DBPath       string
    BackupDir    string
    MaxBackups   int
    CompressData bool
}

type Backup struct {
    config BackupConfig
}

func NewBackup(config BackupConfig) *Backup {
    return &Backup{
        config: config,
    }
}

func (b *Backup) CreateBackup() error {
    // 确保备份目录存在
    if err := os.MkdirAll(b.config.BackupDir, 0750); err != nil {
        return fmt.Errorf("failed to create backup directory: %w", err)
    }

    timestamp := time.Now().Format("20060102_150405")
    backupFile := filepath.Join(b.config.BackupDir, fmt.Sprintf("dashboard_%s.db", timestamp))

    // 确保数据库写入已完成
    db, err := os.Open(b.config.DBPath)
    if err != nil {
        return fmt.Errorf("failed to open database: %w", err)
    }
    db.Close()

    // 复制数据库文件
    if err := copyFile(b.config.DBPath, backupFile); err != nil {
        return fmt.Errorf("failed to copy database: %w", err)
    }

    // 压缩备份
    if b.config.CompressData {
        if err := compressFile(backupFile); err != nil {
            return fmt.Errorf("failed to compress backup: %w", err)
        }
        // 删除未压缩的文件
        os.Remove(backupFile)
        backupFile += ".gz"
    }

    // 清理旧备份
    return b.cleanOldBackups()
}

func (b *Backup) cleanOldBackups() error {
    files, err := filepath.Glob(filepath.Join(b.config.BackupDir, "dashboard_*.db*"))
    if err != nil {
        return err
    }

    if len(files) <= b.config.MaxBackups {
        return nil
    }

    // 按修改时间排序
    type fileInfo struct {
        path string
        time time.Time
    }
    
    var filesInfo []fileInfo
    for _, file := range files {
        info, err := os.Stat(file)
        if err != nil {
            continue
        }
        filesInfo = append(filesInfo, fileInfo{file, info.ModTime()})
    }

    // 按时间排序
    sort.Slice(filesInfo, func(i, j int) bool {
        return filesInfo[i].time.Before(filesInfo[j].time)
    })

    // 删除最旧的文件
    for i := 0; i < len(filesInfo)-b.config.MaxBackups; i++ {
        os.Remove(filesInfo[i].path)
    }

    return nil
}

func copyFile(src, dst string) error {
    input, err := os.ReadFile(src)
    if err != nil {
        return err
    }
    return os.WriteFile(dst, input, 0640)
}

func compressFile(file string) error {
    cmd := exec.Command("gzip", file)
    return cmd.Run()
} 