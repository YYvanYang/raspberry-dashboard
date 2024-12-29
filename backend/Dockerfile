FROM golang:1.22-alpine AS builder

# 添加构建依赖
RUN apk add --no-cache gcc musl-dev

# 设置构建缓存
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/go-mod-cache

WORKDIR /build
COPY go.* ./
RUN --mount=type=cache,target=/go-mod-cache \
    go mod download

COPY . .
RUN --mount=type=cache,target=/go-cache \
    --mount=type=cache,target=/go-mod-cache \
    CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o main ./cmd

FROM alpine:latest

# 添加运行时依赖
RUN apk add --no-cache sqlite-libs tzdata ca-certificates curl

# 创建非 root 用户
RUN adduser -D -u 1000 appuser

WORKDIR /app
COPY --from=builder /build/main .
COPY --from=builder /build/configs ./configs

# 设置文件权限
RUN chown -R appuser:appuser /app && \
    chmod -R 550 /app && \
    chmod -R 660 /app/configs

# 切换到非 root 用户
USER appuser

EXPOSE 3001
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:3001/health || exit 1

ENTRYPOINT ["./main"] 