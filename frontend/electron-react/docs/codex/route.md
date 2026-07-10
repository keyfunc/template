# 路由规范

## 适用范围

- `src/routes/`：TanStack Router 文件路由、根布局与页面挂载

## 路由约束

- 使用 `createFileRoute` / `createRootRoute`，路由对象统一导出为 `Route`
- 路由文件只声明路径、加载边界、布局和 `src/page/` 页面挂载，不堆放具体业务实现
- 根级 Provider 放在明确的应用入口或根路由中，不在普通页面重复创建
- 文件路由命名遵循 TanStack Router 约定，不为迁就页面目录自行改变路由结构

## Electron 适配

- `src/main.tsx` 当前使用 Hash History，`vite.config.ts` 当前使用相对 `base`，用于适配 Electron 本地文件加载
- 修改 History、`base`、路由入口或代码分割配置后，必须验证 Web 构建和 Electron 中的首次加载、跳转与刷新
- `src/routeTree.gen.ts` 是自动生成文件，禁止手动编辑；需要调整时修改 `src/routes/` 或生成配置

页面职责见 [page.md](./page.md)，Electron 加载规则见 [electron.md](./electron.md)。
