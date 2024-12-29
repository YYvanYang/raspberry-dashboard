FROM node:20-alpine AS builder

# 设置构建参数
ARG API_URL
ENV VITE_API_URL=${API_URL}

WORKDIR /app

# 安装依赖
COPY package*.json ./
RUN --mount=type=cache,target=/root/.npm \
    npm ci

# 构建应用
COPY . .
RUN npm run build

FROM node:20-alpine

# 添加运行时依赖
RUN apk add --no-cache tzdata curl

# 创建非 root 用户
RUN adduser -D -u 1000 appuser

WORKDIR /app

# 只复制必要的文件
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/build ./build

# 只安装生产依赖
RUN --mount=type=cache,target=/root/.npm \
    npm ci --omit=dev

# 设置文件权限
RUN chown -R appuser:appuser /app

# 切换到非 root 用户
USER appuser

# 设置环境变量
ENV NODE_ENV=production
ENV PORT=3000
ENV ORIGIN=http://localhost:3000

EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:3000/health || exit 1

CMD ["node", "build"] 