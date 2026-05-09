# AGENTS

## 项目说明

vue3/ts/tailwindcss 模版

## 全局约束

- 包管理使用 `pnpm`，业务逻辑优先 TypeScript
- 代码风格以项目现有 ESLint / Prettier 为准；默认遵循 2 空格缩进、单引号、不使用分号、行宽 100、命名 PascalCase
- 封装业务函数补充单行注释：`// 注释说明`
- `props`、`emits`、`defineExpose`、类型别名等对外接口需要注释；字段注释使用 `/** 注释内容 */`
- 能直接复用现有类型时，不额外声明等价类型别名；仅在需要收窄、组合、扩展语义或提升可读性时新增
- 动态属性、组件 `props` 绑定及相关字段统一使用 camelCase，禁止 kebab-case
- 样式优先使用 TailwindCSS 工具类，布局优先 `flex`；主题色、字体大小、背景色等可复用设计值优先沉淀为 CSS 变量
- 新建页面与组件时使用 PascalCase 文件夹承载模块，入口文件统一命名为 `index.vue`，配套类型、样式、常量等文件放在同级目录
- 不要加任何的兜底代码，防御性代码，该是什么就使用什么，能一个函数写完，就不要拆分多个函数

## 分层职责

- `src/pages/` 负责页面编排与业务流程落点
- `src/components/` 负责通用组件与交互封装
- `src/hooks/` 负责可复用组合式逻辑
- `src/services/` 负责接口请求、参数转换与数据标准化
- `src/stores/` 负责跨页面共享状态
- `src/utils/` 负责基础工具能力
- `src/assets/` 负责图片、字体、音视频等静态资源

## 覆盖规则

- 目录级约束写在对应目录的 `AGENTS.md`
- 若子目录存在 `AGENTS.md`，以子目录规则为准
