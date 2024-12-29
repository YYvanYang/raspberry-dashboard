#!/bin/bash

echo "启动开发环境..."

# 启动后端服务
cd backend
go run cmd/main.go &
BACKEND_PID=$!

# 启动前端服务
cd ../frontend
npm run dev &
FRONTEND_PID=$!

# 输出访问地址
echo -e "\n服务已启动:"
echo "前端服务运行在: http://localhost:5173"
echo "后端服务运行在: http://localhost:3001"
echo "按 Ctrl+C 停止服务"

# 捕获 CTRL+C 信号
trap 'echo -e "\n停止服务..."; kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; exit' INT TERM

# 等待任意子进程结束
wait 