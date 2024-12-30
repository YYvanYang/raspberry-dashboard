#!/bin/bash

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# 获取项目根目录
PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

echo "初始化开发环境..."

# 初始化前端
cd "$PROJECT_ROOT/frontend"
echo "Installing frontend dependencies..."
pnpm install

# 初始化后端
cd "$PROJECT_ROOT/backend"
echo "Initializing backend..."
./scripts/dev-init.sh

# 复制环境变量文件(如果不存在)
if [ ! -f ".env.development" ]; then
    echo "Creating development environment file..."
    cp .env.example .env.development
    echo "请编辑 backend/.env.development 配置环境变量"
fi

echo "开发环境初始化完成！"
echo "启动开发服务器:"
echo "1. 前端: cd frontend && pnpm dev"
echo "2. 后端: "
echo "   - 配置环境变量: 编辑 backend/.env.development"
echo "   - 启动服务: cd backend && go run cmd/main.go"
echo ""
echo "或者使用: ./scripts/dev.sh 同时启动前后端服务"