#!/bin/bash

echo "启动开发环境..."

# 加载后端环境变量
if [ -f "backend/.env.development" ]; then
    export $(cat backend/.env.development | xargs)
fi

# 启动前端服务
cd frontend && pnpm dev &
FRONTEND_PID=$!

# 启动后端服务
cd ../backend && go run cmd/main.go &
BACKEND_PID=$!

echo -e "\n服务已启动:"
echo "前端服务运行在: http://localhost:5173"
echo "后端服务运行在: http://localhost:3001"
echo "按 Ctrl+C 停止服务"

# 捕获 SIGINT 信号（Ctrl+C）
trap "kill $FRONTEND_PID $BACKEND_PID; exit" INT

# 等待任意子进程结束
wait 