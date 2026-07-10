# 状态模块规范

## 适用范围

- `src/store/`：Zustand 全局客户端状态

## 目录与命名

- 一个稳定业务域对应一个 Store 文件，推荐路径为 `src/store/<domain>.ts`
- Store Hook 使用 `useXxxStore`，State、Action 与公开字段使用清晰业务名称
- Store State、Action 参数和关键字段补充简短中文注释

## 状态边界

- Store 只保存跨页面共享的客户端状态，例如鉴权信息、用户偏好和桌面 UI 偏好
- 服务端数据、加载状态和请求错误交给 TanStack Query，不复制到 Zustand
- 组件私有状态保留在组件内，不因可能复用就提前提升为全局状态
- 派生值优先通过选择器计算，不同时保存可由现有状态推导出的重复数据

## 持久化

- 需要持久化时使用 Zustand middleware，并使用稳定、具备业务含义的存储键
- 持久化结构变化时明确迁移策略，不通过随意默认值掩盖不兼容数据
- Store 不直接承担接口请求和响应标准化；相关逻辑放在 `src/service/` 或 Query Hook

相关规范见 [hook.md](./hook.md)、[service.md](./service.md) 与 [api.md](./api.md)。
