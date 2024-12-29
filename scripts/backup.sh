#!/bin/bash

# 备份配置
BACKUP_DIR="/backup/db"
DATA_DIR="/app/data"
MAX_BACKUPS=7
DATE=$(date +%Y%m%d_%H%M%S)

# 创建备份目录
mkdir -p $BACKUP_DIR

# 等待数据库解锁
while [ -f $DATA_DIR/dashboard.db-wal ]; do
    sleep 1
done

# 创建备份
cp $DATA_DIR/dashboard.db $BACKUP_DIR/dashboard_${DATE}.db
gzip $BACKUP_DIR/dashboard_${DATE}.db

# 清理旧备份
find $BACKUP_DIR -name "dashboard_*.db.gz" -mtime +$MAX_BACKUPS -delete

# 检查备份是否成功
if [ $? -eq 0 ]; then
    echo "Backup completed successfully: dashboard_${DATE}.db.gz"
else
    echo "Backup failed!"
    exit 1
fi 