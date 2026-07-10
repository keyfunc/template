# Electron 规范

## 进程边界

- `script/main.ts` 是当前主进程入口，负责窗口、应用生命周期和桌面系统能力
- `src/` 属于 Web 渲染进程，不直接导入主进程模块或启用 Node.js 集成
- 渲染进程需要调用原生能力时定义最小 preload/IPC 接口，保持 `contextIsolation`，不开启 `nodeIntegration`；纯主进程能力不额外暴露
- IPC 参数和返回值使用可序列化的明确类型；主进程校验来自渲染进程的参数和调用权限

## 窗口与布局

- BrowserWindow 尺寸变化会自动改变渲染层 viewport，响应式布局由 CSS Grid、Flex、断点和容器查询完成
- 根据最窄可用布局设置合理的 `minWidth`、`minHeight`，需要按内容区计算时使用 `useContentSize`
- 不通过 `webContents` 缩放或动态根字号代替响应式布局
- 调整标题栏、窗口边框或内容尺寸后，检查各平台实际可用内容区域

## 开发与打包

- 当前 `pnpm dev` 只运行 Vite；`pnpm dev:app` 通过 `script/main.ts` 加载已有的 `dist/index.html`
- 修改渲染层后先执行 `pnpm build`，再使用当前桌面启动流程验证
- `electron-builder.yaml` 的产物目录是 `release/`；修改入口、平台目标、应用图标或打包资源时验证相应平台产物
- 路由使用 Hash History 且 Vite 使用相对 `base`；修改加载地址或开发流程时同时验证开发与打包环境
- 窗口、IPC、生命周期和打包配置属于公共基础能力，修改前确认影响范围，修改后至少运行 `pnpm build`

布局细则见 [style.md](./style.md)，路由与本地加载细则见 [route.md](./route.md)。
