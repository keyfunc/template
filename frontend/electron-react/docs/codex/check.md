# 自检清单

## 修改前

- 已读取根目录 `AGENTS.md` 和本次改动对应的 `docs/codex/*.md`
- 已确认目标职责与目录名称符合本项目固定结构
- 已检查工作区现有改动，不覆盖或重排无关内容
- 已确认目标不是 `src/routeTree.gen.ts` 等生成文件

## 代码自检

- 业务逻辑使用 TypeScript，未新增 `any` 或等价类型
- 新增业务函数、公共类型、Props、Store State、请求参数和响应类型已补充必要中文注释
- 页面负责流程编排，可复用能力已放入对应的 `component`、`hook`、`service`、`store` 或 `util`
- React 组件未因习惯性优化添加 `useMemo` / `useCallback`
- 未新增无业务依据的兜底、防御性分支或静默吞错

## 样式自检

- 应用外壳及主要分区按需优先使用 Grid，局部排列使用 Flex/Grid，并通过断点或容器查询适配可用空间
- 未使用动态根字号让 Electron 页面整体等比缩放
- 弹性区域已检查 `minmax(0, 1fr)`、`min-w-0`、`min-h-0` 和滚动归属
- 纯布局变化由 CSS 完成，窗口尺寸未写入 Zustand
- 公共设计值使用语义化 CSS 变量，页面私有样式未放入 `src/index.css`

## 数据与 Electron 自检

- 页面和组件未直接创建 Axios 实例或处理通用请求细节
- 新接口位于 `src/service/` 对应业务域，服务端状态由 TanStack Query 管理
- Zustand 未保存可由 Query 获取的服务端缓存
- 主进程与渲染进程职责清晰，原生能力未通过不安全方式直接暴露给页面
- 路由历史、Vite `base`、BrowserWindow 或打包配置改动已考虑本地文件加载

## 提交前

- 至少执行 `pnpm lint`
- 涉及 TypeScript、路由、Vite、Electron 主进程或公共请求层时执行 `pnpm build`
- 涉及 electron-builder、桌面资源或打包入口时，在具备对应工具链的环境验证受影响目标；全平台脚本只在环境支持时运行
- 提交信息遵循 Conventional Commits；未在未经确认的情况下运行会执行 `git add .` 的 `pnpm commit`
