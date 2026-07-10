# 桌面端样式与响应式规范

## 基础约束

- 样式优先使用 TailwindCSS 4 工具类，不引入额外运行时 className 或 CSS-in-JS 方案
- `src/index.css` 只维护 Tailwind 入口、全局基础样式和稳定设计变量，不存放页面私有样式
- 应用外壳与主要分区优先使用 CSS Grid，工具栏、表单行和组件内部排列使用 Flex 或 Grid
- `className` 保持可读；复杂条件类先整理为语义清晰的局部变量或小型工具函数

## 桌面布局

- 应用外壳通常占满 `h-dvh`，使用明确的行列划分组织顶部栏、侧栏、内容区和状态栏
- 弹性列使用 `minmax(0, 1fr)`；Flex/Grid 子项按需添加 `min-w-0` 与 `min-h-0`，防止内容撑破窗口
- 明确每个区域的滚动职责，需要滚动的内容区使用 `overflow-auto`，避免窗口和内部面板同时滚动
- 侧栏或详情栏可使用 `clamp()` 限制伸缩范围，主内容区使用 `fr` 或百分比占据剩余空间
- 阅读型内容可以设置合理的最大宽度；表格、画布和工作区按业务需要使用可用空间

推荐的应用外壳形式：

```tsx
<div className="grid h-dvh grid-rows-[auto_minmax(0,1fr)_auto]">
	<header>顶部栏</header>
	<div className="grid min-h-0 grid-cols-[clamp(14rem,18vw,18rem)_minmax(0,1fr)] max-lg:grid-cols-1">
		<aside className="overflow-auto max-lg:hidden">侧栏</aside>
		<main className="min-h-0 min-w-0 overflow-auto">内容区</main>
	</div>
	<footer>状态栏</footer>
</div>
```

## 响应式策略

- 断点依据内容何时无法正常显示来确定，不按电脑、平板或手机等设备名称判断
- 宽窗口可以展示多栏；空间不足时折叠侧栏、把详情栏改为抽屉，或隐藏次要信息
- 可复用组件优先根据自身容器使用容器查询，页面级结构再使用媒体查询或 Tailwind 断点
- 纯布局变化交给 CSS；只有业务行为依赖尺寸时才使用 `matchMedia` 或 `ResizeObserver`
- 不监听 `resize` 后把窗口宽高存入 Zustand，不使用 `webContents` 缩放代替响应式布局
- Electron 窗口的 `minWidth`、`minHeight` 和可选 `useContentSize` 应与最窄可用布局保持一致

## 尺寸与变量

- 不动态修改根字号让界面随窗口整体等比缩放；`rem` 用于字号和间距体系
- 细边框和小图标可以使用 `px`；弹性尺寸使用 `fr`、百分比与 `minmax()`
- 需要有限连续变化时使用 `clamp()`，同时设置符合业务可用性的上下限
- 主题色、文字色、背景色、字号等稳定值以语义化 CSS 变量维护，避免按视觉结果命名
- 暗色模式、缩放体系和新断点只在需求明确时建立，不预先添加无业务依据的结构
