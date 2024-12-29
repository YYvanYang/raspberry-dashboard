#!/bin/bash

# 设置变量
RASPBERRY_PI_HOST="your-raspberry-pi-host"  # 替换为您的树莓派地址
RASPBERRY_PI_USER="pi"                      # 替换为您的树莓派用户名
IMAGE_NAME="raspberry-dashboard"
IMAGE_TAG="latest"
CONTAINER_NAME="dashboard"

# 保存镜像
echo "Saving docker image..."
docker save ${IMAGE_NAME}:${IMAGE_TAG} | gzip > dashboard.tar.gz

# 传输到树莓派
echo "Transferring to Raspberry Pi..."
scp dashboard.tar.gz ${RASPBERRY_PI_USER}@${RASPBERRY_PI_HOST}:~/

# 在树莓派上执行部署
echo "Deploying on Raspberry Pi..."
ssh ${RASPBERRY_PI_USER}@${RASPBERRY_PI_HOST} << 'EOF'
    # 加载镜像
    gunzip -c dashboard.tar.gz | docker load

    # 停止并删除旧容器（如果存在）
    docker stop ${CONTAINER_NAME} 2>/dev/null || true
    docker rm ${CONTAINER_NAME} 2>/dev/null || true

    # 运行新容器
    docker run -d \
        --name ${CONTAINER_NAME} \
        --restart unless-stopped \
        -p 3001:3001 \
        -v /data/dashboard:/app/data \
        -v /data/dashboard/logs:/app/logs \
        ${IMAGE_NAME}:${IMAGE_TAG}

    # 清理
    rm dashboard.tar.gz
EOF

# 清理本地文件
rm dashboard.tar.gz

echo "Deployment completed successfully!" 