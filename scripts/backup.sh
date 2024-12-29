#!/bin/bash

# 设置变量
RASPBERRY_PI_HOST="your-raspberry-pi-host"  # 替换为您的树莓派地址
RASPBERRY_PI_USER="pi"                      # 替换为您的树莓派用户名
BACKUP_DIR="backups"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

# 创建本地备份目录
mkdir -p ${BACKUP_DIR}

# 在树莓派上执行备份
echo "Creating backup on Raspberry Pi..."
ssh ${RASPBERRY_PI_USER}@${RASPBERRY_PI_HOST} << 'EOF'
    cd /data/dashboard
    tar czf dashboard_backup_${TIMESTAMP}.tar.gz data/
EOF

# 下载备份文件
echo "Downloading backup..."
scp ${RASPBERRY_PI_USER}@${RASPBERRY_PI_HOST}:/data/dashboard/dashboard_backup_${TIMESTAMP}.tar.gz ${BACKUP_DIR}/

# 清理远程备份文件
ssh ${RASPBERRY_PI_USER}@${RASPBERRY_PI_HOST} "rm /data/dashboard/dashboard_backup_${TIMESTAMP}.tar.gz"

echo "Backup completed successfully! Saved to ${BACKUP_DIR}/dashboard_backup_${TIMESTAMP}.tar.gz" 