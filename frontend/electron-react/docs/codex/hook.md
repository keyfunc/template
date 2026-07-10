# Hook 规范

## 适用范围

- `src/hook/`：跨页面或跨组件复用的 React Hook、Query Hook 与副作用组合

## 目录与命名

- Hook 文件和导出统一使用 `useXxx`，推荐路径为 `src/hook/useXxx.ts`
- Hook 聚焦单一职责；仅由一个页面或组件使用的逻辑优先与调用方同目录维护
- 对外参数、返回值和回调 payload 使用明确类型；复杂共享类型再拆到相邻 `type.ts`

## 状态与副作用

- 服务端数据通过 TanStack Query 管理；Query Hook 使用带业务域前缀的稳定数组作为 query key
- 不把 Query 缓存复制进 Zustand，也不在 Hook 内重复实现请求层已有的通用错误处理
- 浏览器事件、Electron 订阅、定时器和观察器等副作用必须成对初始化与清理
- 依赖项表达真实数据关系，不通过关闭规则或省略依赖掩盖生命周期问题
- 仅当引用稳定性属于调用契约或有性能证据时使用 `useMemo` / `useCallback`

## 依赖方向

- Hook 可以依赖 `src/service/`、`src/store/` 与 `src/util/`
- `src/service/` 和 `src/util/` 不反向依赖 Hook、页面或组件
- 页面和组件通过 Hook 消费可复用查询、状态和副作用逻辑

相关规范见 [service.md](./service.md)、[store.md](./store.md) 与 [util.md](./util.md)。
