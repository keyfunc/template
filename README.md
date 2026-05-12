# Template

一个用于沉淀前端、后端、移动端和基础服务脚手架的模板集合。根目录的 `templates.json` 记录了当前可用模板及其路径。

## 推荐创建方式

推荐使用 `create-clear` npm 包创建 template 项目，并按命令行提示选择需要的模板。

## 项目列表

| 项目 | 路径 | 用途 | 技术选型 |
| --- | --- | --- | --- |
| react | `frontend/react` | 通用 React Web 应用模板 | React 19、TypeScript、Vite、TanStack Router、TanStack Query、Tailwind CSS、Axios、Zustand、Orval、Biome |
| vue | `frontend/vue` | 通用 Vue H5/Web 应用模板 | Vue 3、TypeScript、Vite、Vue Router、Pinia、TanStack Query、Tailwind CSS、Axios、ESLint、Prettier |
| admin-react | `frontend/admin-react` | 管理后台前端模板 | React 19、TypeScript、Vite、TanStack Router、TanStack Query、Radix UI、shadcn/ui、Tailwind CSS、Zustand、i18next、Orval、Biome |
| mobile-taro | `frontend/mobile-taro` | 多端移动应用模板，面向小程序和 H5 | Taro v4、React、TypeScript、Tailwind CSS、Axios、Zustand、ESLint、Stylelint |
| go-standard | `backend/go-standard` | 基于 Go 标准库的轻量 API 服务模板 | Go、net/http、PostgreSQL、pgx、YAML 配置、OpenAPI |
| admin-gin | `backend/admin-gin` | 管理后台 API 服务模板 | Go、Gin、GORM、PostgreSQL、Redis、JWT、Viper、slog、OpenAPI |
| devops/docker | `devops/docker` | 本地开发基础服务配置 | Docker Compose、MySQL、PostgreSQL、Redis |

## 说明

- 每个子模板可独立使用，具体细节以子目录 README 或配置文件为准。
- 后端 OpenAPI 文件可供前端 Orval 生成接口请求代码。
- 项目采用 Apache License 2.0。
