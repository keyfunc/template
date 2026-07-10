# 页面编排规范

## 适用范围

- `src/page/`：页面级视图、业务流程和模块编排

## 目录规范

- 页面目录使用 PascalCase，入口为 `src/page/<PageName>/index.tsx`
- `src/routes/` 只声明路由、布局和页面挂载；具体页面实现放在 `src/page/`
- 页面负责组装 `src/component/`、`src/hook/`、`src/store/` 与 `src/service/`，不沉淀通用能力
- 页面私有组件和逻辑可保留在页面目录；确认可跨页面复用后再移动到公共分层

## 实现约束

- 页面事件处理函数使用明确动词语义，例如 `handleSubmit`、`handleOpenDetail`
- 请求 URL、请求头、响应解析和数据转换不直接写在页面中
- 页面只处理当前业务流程相关的加载、空状态、错误交互、跳转和弹窗
- 演示页、联调页和测试页需要明确用途，避免临时代码直接成为正式实现
- 页面根容器明确高度、最小宽高与滚动归属，避免页面和内部区域同时滚动

相关规范见 [component.md](./component.md)、[hook.md](./hook.md)、[service.md](./service.md)、[route.md](./route.md) 与 [style.md](./style.md)。
