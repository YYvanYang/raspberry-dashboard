#!/bin/bash

# 备份目录
BACKUP_DIR="backups"
mkdir -p "$BACKUP_DIR"

# 生成备份文件名
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/backup_$TIMESTAMP.tar.gz"

# 停止服务
docker compose down

# 创建备份
tar -czf "$BACKUP_FILE" data config .env

# 重启服务
docker compose up -d

# 删除7天前的备份
find "$BACKUP_DIR" -name "backup_*.tar.gz" -mtime +7 -delete

echo "Backup created: $BACKUP_FILE" 