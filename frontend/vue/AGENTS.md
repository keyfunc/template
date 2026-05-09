# AGENTS

## 项目说明

Vue 3 / TypeScript / Vite 8 / TailwindCSS 4 模板项目。路由使用 Vue Router，状态使用 Pinia，服务端状态使用 TanStack Query，请求层使用 Axios。

## 核心约束

- 包管理使用 `pnpm`；常用命令：`pnpm dev`、`pnpm build`、`pnpm lint`、`pnpm format`
- 代码风格以 ESLint / Prettier 为准：2 空格、单引号、无分号、行宽 100、尾随逗号
- `src` 内导入优先使用 `@/` 别名；业务代码优先 TypeScript，避免 `any`
- 新增业务函数、公共类型、组件 props / emits / expose、请求参数和响应类型时，补充简短中文注释
- 不添加无业务依据的兜底代码、防御性代码或静默吞错；接口契约不清晰时先补类型或明确调用方约束
- 修改请求拦截器、路由入口、状态入口、构建配置等公共能力时，运行必要校验

## Vue 约定

- 组件使用 `.vue`，有逻辑时优先 `<script setup lang="ts">`
- 组件名使用 PascalCase，组合式函数命名为 `useXxx`，Pinia store 命名为 `useXxxStore`
- 动态属性、组件 props 绑定和 TS 字段统一使用 camelCase
- 页面放在 `src/page/<route-name>/index.vue`；通用组件放在 `src/component/<ComponentName>/index.vue`
- 接口请求统一走 `src/util/request.ts`，业务接口按领域放在 `src/service/<domain>/index.ts`
- 服务端状态优先使用 TanStack Query；Pinia 只存跨页面共享的客户端状态
- 样式优先 TailwindCSS 工具类，布局优先 `flex` / `grid`，复杂复用设计值沉淀为 CSS 变量

## 目录职责

- `src/router/`：路由声明和页面挂载
- `src/page/`：页面编排与业务流程
- `src/component/`：通用 Vue 组件；不存在时按需新建
- `src/hook/`：可复用组合式逻辑和 Query hook；不存在时按需新建
- `src/service/`：接口请求、数据转换、TanStack Query 基础配置
- `src/store/`：Pinia 全局客户端状态
- `src/util/`：基础工具和通用基础设施
- `public/`：无需打包处理的静态资源

## 提交与覆盖

- 提交信息遵循 Conventional Commits，项目已配置 commitlint 和 commitizen
- 提交前至少运行 `pnpm lint`；涉及类型、路由、构建配置或公共请求层时运行 `pnpm build`
- 子目录存在 `AGENTS.md` 时，以子目录规则为准
