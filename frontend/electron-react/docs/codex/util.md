# 工具模块规范

## 适用范围

- `src/util/`：请求实例及与具体页面无关的通用基础工具

## 目录规范

- 按能力拆分子目录，沿用当前 `request/`、`date/`、`number/` 等单数目录命名
- 通用工具推荐入口为 `src/util/<capability>/index.ts`，不要在 `src/util/` 根目录堆叠无关文件
- 函数和目录使用准确语义，避免 `helper`、`common`、`commonUtil` 等模糊名称

## 职责边界

- 工具函数优先保持纯函数特性，不隐式修改外部状态
- 工具层不依赖页面、组件、Hook 或 Store；需要平台 API 时把边界和副作用写清楚
- 已明显具备业务域语义的逻辑放入 `src/service/` 或对应页面模块，不为复用而强行下沉
- 日期、数字等能力优先复用现有依赖或同目录实现，不重复声明等价工具
- 请求基础能力集中在 `src/util/request/`，业务接口仍放在 `src/service/`

相关规范见 [api.md](./api.md)、[service.md](./service.md) 与 [hook.md](./hook.md)。
