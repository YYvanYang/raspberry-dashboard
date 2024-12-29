#!/bin/bash

# 设置 GOOS 以支持在 WSL2 中编译 Windows 程序
export GOOS=linux

# 启动后端服务
cd backend
go run cmd/main.go &
BACKEND_PID=$!

# 启动前端服务
cd ../frontend
# 使用 npm 的跨平台启动
npm run dev -- --host 0.0.0.0 &
FRONTEND_PID=$!

# 捕获 CTRL+C 信号
trap 'kill $BACKEND_PID $FRONTEND_PID 2>/dev/null' INT TERM

# 输出访问地址
echo "前端服务运行在: http://localhost:5173"
echo "后端服务运行在: http://localhost:3001"
echo "按 Ctrl+C 停止服务"

# 等待任意子进程结束
wait 