# electron-react

基于 Electron、React 和 TypeScript 的桌面应用项目。渲染层由 Vite 构建，使用 TanStack Router 管理路由、TanStack Query 管理服务端状态、Zustand 管理全局客户端状态，并通过 Axios 访问接口。

## 技术栈

- Electron 43 / electron-builder
- React 19 / React Compiler
- TypeScript 6 / Vite 8
- TailwindCSS 4
- TanStack Router / TanStack Query
- Zustand / Axios
- Biome / Husky / commitlint

## 开始使用

安装依赖：

```bash
pnpm install
```

启动渲染层开发服务器：

```bash
pnpm dev
```

当前 `pnpm dev` 只启动 Vite。Electron 主进程会加载已有的 `dist/index.html`，因此本地运行桌面窗口前需要先构建渲染层：

```bash
pnpm build
pnpm dev:app
```

打包桌面应用：

```bash
pnpm build
pnpm build:app
```

安装包输出到 `release/`。`build:app` 当前配置了 Windows NSIS、macOS DMG 和 Linux AppImage，实际可构建目标受运行平台与本机构建环境限制。

## 常用命令

| 命令 | 说明 |
| --- | --- |
| `pnpm dev` | 启动 Vite 开发服务器 |
| `pnpm build` | TypeScript 校验并构建渲染层 |
| `pnpm preview` | 预览 Vite 构建结果 |
| `pnpm dev:app` | 使用已有 `dist/` 启动 Electron |
| `pnpm build:app` | 使用 electron-builder 打包应用 |
| `pnpm lint` | 使用 Biome 检查项目 |
| `pnpm format` | 使用 Biome 格式化项目 |
| `pnpm commit` | 暂存当前改动并进入交互式提交 |

## 目录结构

```text
electron-react/
├── script/                 # Electron 主进程代码
├── electron/               # 预留的桌面端打包资源
├── public/                 # 构建时原样复制的公开资源
├── src/
│   ├── assets/             # 由 Vite 处理的图片、字体等资源
│   ├── component/          # 可复用 React 组件
│   ├── hook/               # 可复用 Hook 与 Query Hook
│   ├── page/               # 页面编排与业务流程
│   ├── routes/             # TanStack Router 文件路由
│   ├── service/            # 按业务域组织的接口服务
│   ├── store/              # Zustand 全局客户端状态
│   ├── util/               # 请求实例与通用工具
│   ├── index.css           # Tailwind 入口与全局设计变量
│   ├── main.tsx            # React 渲染入口
│   └── routeTree.gen.ts    # TanStack Router 自动生成文件
├── docs/codex/             # 按模块拆分的开发规范
├── electron-builder.yaml   # 桌面应用打包配置
└── vite.config.ts          # 渲染层构建配置
```

以上目录名是项目约定：除 `routes`、`assets` 外，业务分层目录使用当前确定的单数形式，不改写为 `components`、`hooks`、`pages`、`services`、`stores` 或 `utils`。

## 架构约定

- `script/` 与 `src/` 分别属于 Electron 主进程和 Web 渲染进程；渲染层需要系统能力时通过明确的 preload/IPC 边界接入。
- `src/routes/` 只声明路由、布局和页面挂载，页面实现放在 `src/page/`；`src/routeTree.gen.ts` 禁止手动编辑。
- 业务请求统一复用 `src/util/request/`，接口调用与数据标准化放在 `src/service/`，页面或 Query Hook 负责消费。
- TanStack Query 保存服务端状态，Zustand 只保存跨页面共享的客户端状态。
- 桌面布局使用 Grid 组织应用外壳、Flex/Grid 组织局部内容，并按可用空间设置断点；不使用动态根字号让整个界面随窗口等比缩放。

完整协作约束见 [AGENTS.md](AGENTS.md)，模块细则见 [docs/codex/](docs/codex/)。
