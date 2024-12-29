# 前端构建阶段
FROM --platform=$BUILDPLATFORM node:20-alpine AS frontend-builder
WORKDIR /app
COPY frontend/ .
RUN npm ci && npm run build

# 后端构建阶段
FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS backend-builder
WORKDIR /app
# 设置 Go 编译缓存
RUN --mount=type=cache,target=/root/.cache/go-build \
    go env -w GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/go/pkg/mod \
    go env -w GOMODCACHE=/go/pkg/mod

# 复制后端代码和前端构建产物
COPY backend/ .
COPY --from=frontend-builder /app/build ./web

# 下载依赖
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 交叉编译为 ARM64（树莓派）
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
RUN go build -ldflags="-w -s" -o server ./cmd/main.go

# 最终运行镜像
FROM --platform=linux/arm64 alpine:latest

# 安装必要的运行时依赖
RUN apk add --no-cache tzdata ca-certificates

# 创建非 root 用户
RUN adduser -D -u 1000 appuser

WORKDIR /app

# 复制构建产物
COPY --from=backend-builder /app/server .
COPY --from=backend-builder /app/web ./web

# 创建数据和日志目录
RUN mkdir -p /app/data /app/logs && \
    chown -R appuser:appuser /app && \
    chmod -R 755 /app

# 切换到非 root 用户
USER appuser

# 暴露端口
EXPOSE 3001

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:3001/health || exit 1

# 启动服务
CMD ["./server"] 