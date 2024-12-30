#!/bin/bash

# 设置开发环境变量
export GO_ENV=development
export JWT_SECRET=development_secret_key

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

echo "Development environment initialized!" 