# 组件规范

## 适用范围

- `src/component/`：跨页面复用的 React 组件与稳定交互封装

## 目录与命名

- 组件及其目录使用 PascalCase，推荐入口为 `src/component/<ComponentName>/index.tsx`
- 简单组件不强制创建目录或 `type.ts`；类型复杂或需要共享时再拆分
- 组件私有逻辑与组件同目录维护，具备跨组件复用价值的 Hook 才放入 `src/hook/`
- 新增组件前先确认现有组件是否可扩展，避免职责重叠的近似实现

## 封装约束

- 使用函数组件和 TypeScript；组件文件使用 `.tsx`，普通逻辑与类型使用 `.ts`
- Props、回调 payload、公开 ref 能力及相关字段需要清晰类型和简短中文注释
- 自定义 Props 与普通 React DOM 属性使用 camelCase，`aria-*`、`data-*` 保留连字符；事件回调使用 `onXxx`，内部处理函数使用 `handleXxx`
- 组件保持通用，不直接依赖页面路由、临时调试状态或未经标准化的接口响应
- 数据请求细节放在 `src/service/`，可复用查询组合放在 `src/hook/`，组件只消费明确的数据与行为
- 能在一个清晰函数中完成的逻辑不机械拆分；仅在复用、职责隔离或可读性有明确收益时拆分
- 项目启用了 React Compiler，不因习惯性优化添加 `useMemo` 或 `useCallback`

## 样式

- 优先使用 TailwindCSS；复杂条件类先整理为语义清晰的局部变量或小型工具函数
- 组件应适应自身可用空间，不依赖页面外层临时 class；必要时使用容器查询
- 组件专属 CSS 就近维护，稳定的主题值沉淀到 `src/index.css`，不把页面私有样式提升为全局样式

详细样式规则见 [style.md](./style.md)。
