#!/bin/bash

# 生成环境变量文件
if [ ! -f .env ]; then
    echo "Creating .env file..."
    echo "JWT_SECRET=$(openssl rand -base64 32)" > .env
    echo "ADMIN_PASSWORD=admin123" >> .env
fi

# 确保数据目录存在
mkdir -p data config

# 启动服务
docker compose up -d

echo "Services started successfully!"
echo "Frontend: http://localhost:3000"
echo "Backend: http://localhost:3001" 