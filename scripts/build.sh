#!/bin/bash

# 设置变量
IMAGE_NAME="raspberry-dashboard"
IMAGE_TAG="latest"

# 构建多架构镜像
echo "Building multi-arch image..."
docker buildx build \
  --platform linux/arm64 \
  -t ${IMAGE_NAME}:${IMAGE_TAG} \
  --load \
  .

echo "Build completed successfully!" 