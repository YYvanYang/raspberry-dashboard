# Raspberry Dashboard Frontend

基于 SvelteKit 构建的树莓派监控仪表盘前端项目。使用 Svelte 5 的新特性（Runes）和 TailwindCSS 构建。

## 技术栈

- SvelteKit 2.0
- Svelte 5.0 (Runes)
- TailwindCSS
- TypeScript
- Vite
- Docker

## 开发环境

1. 安装依赖：

```bash
npm install
```

2. 启动开发服务器：

```bash
npm run dev

# 或者在新标签页中打开
npm run dev -- --open
```

## 环境变量

创建 `.env` 文件（可以从 `.env.example` 复制）：

```bash
cp .env.example .env
```

主要环境变量：

- `VITE_API_URL`: 后端 API 地址

## 构建

### 本地构建

构建生产版本：

```bash
npm run build
```

预览构建结果：

```bash
npm run preview
```

### Docker 构建

使用 Docker Compose 构建和运行：

```bash
# 构建并启动所有服务
docker compose up -d

# 仅构建前端
docker compose build frontend

# 仅启动前端
docker compose up -d frontend
```

## 项目结构

```
src/
├── lib/            # 可重用的组件和工具
│   ├── components/ # Svelte 组件
│   ├── stores/     # 状态管理
│   └── utils/      # 工具函数
├── routes/         # SvelteKit 路由
└── app.css        # 全局样式
```

## 特性

- 服务端渲染 (SSR)
- 响应式设计
- 实时系统监控
- 日志查看器
- 配置管理
- 用户认证
- 错误边界处理
- 实时通知

## 开发规范

- 组件使用 PascalCase 命名
- 使用 TypeScript 进行类型检查
- 使用 ESLint 和 Prettier 进行代码格式化
- 遵循 TailwindCSS 的工具优先原则

## 测试

运行单元测试：

```bash
npm run test
```

## 代码格式化

```bash
# 格式化代码
npm run format

# 检查代码格式
npm run lint
```
