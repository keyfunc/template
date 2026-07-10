# AGENTS

## 项目上下文

这是一个 Electron 43 / React 19 / TypeScript 6 / Vite 8 桌面应用项目。渲染层使用 TailwindCSS 4、TanStack Router、TanStack Query、Zustand 与 Axios，桌面应用通过 electron-builder 打包。

本文件适用于整个仓库。处理具体模块前，同时读取 `docs/codex/` 中对应规范；若以后新增目录级 `AGENTS.md`，以距离目标文件最近的规则为准。

## 全局约束

- 包管理统一使用 `pnpm`；新增或删除依赖后同步维护 `pnpm-lock.yaml`
- 代码风格以 `biome.json` 为准，不引入 ESLint 或 Prettier；使用 tab 缩进、双引号并保留分号
- 常用命令以 `package.json` 为准：`pnpm dev`、`pnpm build`、`pnpm preview`、`pnpm dev:app`、`pnpm build:app`、`pnpm lint`、`pnpm format`
- 业务代码使用 TypeScript，避免 `any`；能复用已有类型时不声明等价类型
- 导入 `src` 内模块优先使用 `@/` 别名，并保持 Biome 自动整理后的 import 顺序
- 新增业务函数、公共类型、组件 Props、Store State、请求参数和响应类型时，补充简短中文注释；字段使用 `/** 注释内容 */`
- 不添加无业务依据的兜底、防御性分支或静默吞错；契约不明确时先补类型或明确调用方约束
- 不手动编辑生成文件，不顺带重构无关代码；修改公共入口、构建配置或请求基建前先确认影响范围

## 固定目录结构

目录名以当前仓库为唯一标准，不把单数目录改成其他项目常见的复数形式。

| 目录 | 职责 |
| --- | --- |
| `script/` | Electron 主进程代码 |
| `electron/` | 预留的 electron-builder 图标等桌面资源 |
| `src/assets/` | 由 Vite 处理的静态资源 |
| `src/component/` | 可复用 React 组件 |
| `src/hook/` | 可复用 Hook、Query Hook 与副作用组合 |
| `src/page/` | 页面编排与业务流程 |
| `src/routes/` | TanStack Router 路由声明、布局和页面挂载 |
| `src/service/` | 接口调用、参数转换和响应数据标准化 |
| `src/store/` | Zustand 全局客户端状态 |
| `src/util/` | 请求实例和通用基础工具 |
| `public/` | 构建时原样复制的公开资源 |

推荐结构：

- 页面：`src/page/<PageName>/index.tsx`
- 可复用组件：`src/component/<ComponentName>/index.tsx`
- 可复用 Hook：`src/hook/useXxx.ts`
- 业务服务：`src/service/<domain>/index.ts`
- Store：`src/store/<domain>.ts`
- 通用工具：`src/util/<capability>/index.ts`

仅在类型复杂或需要跨文件共享时增加 `type.ts`，不要为了目录形式创建空文件。

## Electron 约定

- `script/main.ts` 是当前主进程入口；`src/main.tsx` 是渲染进程入口，不混用两侧 API
- 渲染进程不直接启用或依赖 Node.js 能力；渲染进程需要原生能力时先设计主进程、preload 与 IPC 的最小安全边界，不开启 `nodeIntegration`
- 当前 `pnpm dev` 只启动 Vite，`pnpm dev:app` 加载已有的 `dist/index.html`；不要把它描述成 Electron 热更新流程
- 渲染层使用 Hash History，Vite 使用相对 `base` 以适配本地文件加载；修改路由历史或资源基址时必须验证打包后的页面加载和刷新
- 窗口最小尺寸应与页面最窄可用布局一致；调整 `BrowserWindow`、生命周期、IPC 或打包配置后运行构建，并按影响验证桌面端
- 不在渲染进程中保存主进程专属对象，也不通过 IPC 传递不可序列化数据

详细规则见 `docs/codex/electron.md`。

## React 与路由

- 页面和组件使用函数组件；渲染文件使用 `.tsx`，普通逻辑和类型使用 `.ts`
- 组件使用 PascalCase，Hook 使用 `useXxx`，Store Hook 使用 `useXxxStore`
- 项目启用了 React Compiler，不因习惯性优化添加 `useMemo` 或 `useCallback`；仅在引用稳定性属于 API 契约或已有性能证据时使用
- `src/routes/` 只负责路由声明、布局和挂载 `src/page/` 页面，不在路由文件堆放具体业务实现
- 路由使用 `createFileRoute` / `createRootRoute`，导出名保持为 `Route`
- `src/routeTree.gen.ts` 由 TanStack Router 生成，禁止手动修改
- 页面只负责编排；可复用 UI 放入 `src/component/`，可复用状态或副作用放入 `src/hook/`

详细规则见 `docs/codex/component.md`、`docs/codex/hook.md`、`docs/codex/page.md` 与 `docs/codex/route.md`。

## 请求与状态

- Axios 实例位于 `src/util/request/index.ts`，业务代码通过 `@/util/request` 复用，不创建新的 Axios 实例
- 接口按领域放在 `src/service/<domain>/index.ts`；组件内不编写 URL、请求头、响应解析等请求细节
- 服务端状态优先使用 TanStack Query；Query Hook 放在 `src/hook/`，query key 使用带领域前缀的稳定数组
- Zustand 只保存跨页面共享的客户端状态，例如鉴权信息和用户偏好；不要复制 TanStack Query 的服务端缓存
- `src/store/auth.ts` 是当前鉴权状态示例，token 是否注入请求头应由真实业务接入方案决定
- 当前项目没有配置 Orval 或 OpenAPI 代码生成，不引用不存在的生成目录和命令；正式接入后再补充约束

详细规则见 `docs/codex/service.md`、`docs/codex/api.md` 与 `docs/codex/store.md`。

## 桌面端样式与响应式

- 样式优先使用 TailwindCSS 工具类；全局入口和设计变量放在 `src/index.css`
- 应用外壳及主要分区优先使用 CSS Grid，工具栏、表单行和组件内部排列使用 Flex 或 Grid
- 断点依据内容可用空间确定，不按设备型号判断；可复用组件优先考虑容器查询
- 不使用 H5 动态根字号方案让整个 Electron 界面按窗口宽度等比缩放
- 字号和间距使用 `rem`，细边框和小图标可使用 `px`，弹性区域使用 `fr`、百分比、`minmax()`，有限连续缩放使用 `clamp()`
- 桌面分区布局明确滚动归属：外壳通常占满 `h-dvh`，弹性子项使用 `min-w-0`、`min-h-0`，需要滚动的区域使用 `overflow-auto`
- 窄窗口通过折叠侧栏、抽屉或减少次要信息切换布局，不单纯缩小全部文字和控件
- 纯布局变化交给 CSS，不监听 `resize` 后把窗口尺寸存入 Zustand；只有业务行为确实依赖尺寸时才使用 `matchMedia` 或 `ResizeObserver`
- 主题色、文字色、背景色和字号等稳定设计值优先沉淀为语义化 CSS 变量；页面私有样式不放入全局文件
- `className` 保持可读，复杂条件样式先整理成语义清晰的局部变量或小型工具函数

详细规则见 `docs/codex/style.md`。

## 静态资源

- 需要参与 Vite 构建、哈希和模块导入的资源放在 `src/assets/`
- 需要保持文件名并按公开路径访问的资源放在 `public/`
- `electron/` 是预留的桌面打包资源目录；接入图标等资源后同步配置 `electron-builder.yaml`
- 新增资源前确认是否可复用现有文件，使用语义化名称，不把脚本或业务配置放入资源目录

详细规则见 `docs/codex/assets.md`。

## 校验与提交

- 提交前至少运行 `pnpm lint`
- 涉及 TypeScript、路由、Vite 配置、Electron 主进程或公共请求层时运行 `pnpm build`
- 修改 `electron-builder.yaml`、桌面资源或打包入口时，在具备对应工具链的环境验证受影响目标；`pnpm build:app` 会请求三个平台，不把单机结果视为全平台验证
- 提交信息遵循 Conventional Commits；lint-staged 会对暂存的 JS、TS、JSON、CSS 和 Markdown 文件执行 Biome 检查与修复
- `pnpm commit` 会执行 `git add .`；只有用户明确要求提交并确认暂存范围时才使用

完整自检清单见 `docs/codex/check.md`。
