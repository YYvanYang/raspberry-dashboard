#!/bin/bash

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# 获取项目根目录
PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

# 检查后端环境变量文件是否存在
if [ ! -f "$PROJECT_ROOT/backend/.env.development" ]; then
    echo "错误: 后端环境变量文件不存在"
    echo "请先复制并配置环境变量:"
    echo "cp backend/.env.example backend/.env.development"
    exit 1
fi

echo "启动开发环境..."

# 设置环境变量
export GO_ENV=development

# 加载后端环境变量(过滤注释和空行)
set -a
source "$PROJECT_ROOT/backend/.env.development"
set +a

# 启动前端服务
cd "$PROJECT_ROOT/frontend" && pnpm dev &
FRONTEND_PID=$!

# 启动后端服务
cd "$PROJECT_ROOT/backend" && go run cmd/main.go &
BACKEND_PID=$!

echo -e "\n服务已启动:"
echo "前端服务运行在: http://localhost:5173"
echo "后端服务运行在: http://localhost:3001"
echo "按 Ctrl+C 停止服务"

# 捕获 SIGINT 信号（Ctrl+C）
trap "kill $FRONTEND_PID $BACKEND_PID; exit" INT

# 等待任意子进程结束
wait