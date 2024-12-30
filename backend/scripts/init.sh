#!/bin/bash

# 设置必要的环境变量
export JWT_SECRET="your-secret-key-here"
export GO111MODULE=on

# 清理和下载依赖
echo "Cleaning module cache..."
go clean -modcache

echo "Downloading dependencies..."
go mod download

echo "Tidying dependencies..."
go mod tidy

echo "Verifying dependencies..."
go mod verify

# 创建必要的目录
echo "Creating necessary directories..."
mkdir -p data logs

# 编译项目
echo "Building project..."
go build -o bin/raspberry-dashboard cmd/main.go

echo "Initialization complete!" 