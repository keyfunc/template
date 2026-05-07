# AGENTS

## 项目说明

React 19 / TypeScript / Vite 8 / TailwindCSS 4 模板项目，路由使用 TanStack Router，服务端状态使用 TanStack Query，全局客户端状态使用 Zustand，请求层使用 Axios。

## 全局约束

- 包管理统一使用 `pnpm`；新增、删除依赖后同步维护 `pnpm-lock.yaml`
- 代码风格以 `biome.json` 为准，不引入 ESLint / Prettier；当前格式为 tab 缩进、双引号、保留分号
- 常用命令：`pnpm dev`、`pnpm build`、`pnpm lint`、`pnpm format`、`pnpm api:generate`、`pnpm api:watch`
- 业务代码优先使用 TypeScript，避免 `any`；能复用已有类型时不要声明等价类型
- 导入 `src` 内模块优先使用 `@/` 别名，保持 Biome 自动整理 import 的结果
- 新增业务函数、公共类型、组件 props、store state、请求参数与响应类型时，补充简短中文注释
- 不添加无业务依据的兜底代码、防御性代码或静默吞错；接口契约不清晰时先补类型或明确调用方约束
- 修改生成文件、依赖配置、请求拦截器、路由入口等公共基础能力时，需要额外确认影响范围并运行校验

## React 约定

- 页面和组件使用函数组件；组件文件使用 `.tsx`，普通逻辑和类型文件使用 `.ts`
- 组件命名使用 PascalCase，hook 命名使用 `useXxx`，store hook 命名使用 `useXxxStore`
- 新建页面放在 `src/page/<PageName>/index.tsx`，页面目录使用 PascalCase
- 新建可复用 hook 放在 `src/hook/useXxx.ts`，只封装可复用状态、请求或副作用逻辑
- 项目启用了 React Compiler，不要为了“习惯性优化”随意添加 `useMemo` / `useCallback`；只有引用稳定性是 API 契约或确有性能证据时再使用
- 组件内不要直接写请求细节；请求封装在 `src/service/`，页面通过 hook 或 React Query 使用

## 路由约定

- TanStack Router 文件路由放在 `src/routes/`，使用 `createFileRoute` / `createRootRoute`
- `src/routeTree.gen.ts` 是 TanStack Router 生成文件，禁止手动编辑
- `src/routes/` 只负责路由声明、布局和页面挂载；具体页面内容放在 `src/page/`
- 路由组件导出名保持为 TanStack Router 约定的 `Route`

## 数据与状态

- 接口请求统一通过 `src/util/request.ts` 导出的 `request` 实例，不在业务代码里直接创建新的 Axios 实例
- 业务接口按领域放在 `src/service/<domain>/index.ts`，负责请求、参数转换和响应数据标准化
- OpenAPI 生成接口使用 Orval，配置入口为 `orval.config.ts`，默认读取 `http://127.0.0.1:8080/openapi.json`
- Orval 生成代码输出到 `src/service/generated/`，禁止手动修改；该目录不做 TS 校验，需要调整生成结果时修改 OpenAPI 契约、Orval 配置或 mutator
- Orval 的请求适配器放在 `src/service/mutator/orval.ts`，生成接口必须复用现有 `request` 实例
- 服务端状态优先使用 TanStack Query；query hook 放在 `src/hook/`，query key 使用稳定数组并带领域前缀
- Zustand 只存放跨页面共享的客户端状态，例如登录 token、用户偏好等；不要把服务端缓存搬进 store
- `src/store/auth.ts` 是模板保留的鉴权状态代码；是否把 token 注入请求头由业务接入阶段决定

## 样式约定

- 样式优先使用 TailwindCSS 工具类，布局优先 `flex` / `grid`
- 全局样式入口是 `src/index.css`，Tailwind 通过 `@import "tailwindcss";` 引入
- 主题色、字号、背景色等可复用设计值优先沉淀为 CSS 变量
- 组件内 className 保持可读，复杂条件样式先整理为清晰变量或小型工具函数

## 分层职责

- `src/routes/` 负责路由声明、根布局、嵌套路由和页面挂载
- `src/page/` 负责页面编排与业务流程落点
- `src/hook/` 负责可复用 React hook、React Query hook 和副作用组合
- `src/service/` 负责接口请求、参数转换与数据标准化
- `src/store/` 负责 Zustand 全局客户端状态
- `src/util/` 负责基础工具能力和通用基础设施
- `src/static/`、`public/` 负责图片、字体、音视频等静态资源

## 提交与校验

- 提交信息遵循 Conventional Commits，项目已配置 commitlint
- 提交前至少运行 `pnpm lint`；涉及类型、路由、构建配置或公共请求层时运行 `pnpm build`
- lint-staged 会对暂存的 JS / TS / JSON / CSS / Markdown 文件执行 `biome check --write`

## 覆盖规则

- 目录级约束写在对应目录的 `AGENTS.md`
- 若子目录存在 `AGENTS.md`，以子目录规则为准
