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

FROM nginx:alpine

# 添加运行时依赖
RUN apk add --no-cache tzdata curl

# 创建非 root 用户
RUN adduser -D -u 1000 appuser

# 配置 Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build /usr/share/nginx/html

# 设置文件权限
RUN chown -R appuser:appuser /usr/share/nginx/html && \
    chmod -R 550 /usr/share/nginx/html && \
    touch /var/run/nginx.pid && \
    chown -R appuser:appuser /var/run/nginx.pid && \
    chown -R appuser:appuser /var/cache/nginx

# 切换到非 root 用户
USER appuser

EXPOSE 80
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:80 || exit 1

CMD ["nginx", "-g", "daemon off;"] 