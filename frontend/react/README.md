# React Template

## API 生成

确认本地后端服务已启动，并能访问 `http://127.0.0.1:8080/openapi.json`，然后执行：

```bash
pnpm api:generate
```

Orval 配置在 `orval.config.ts`，生成代码会输出到 `src/service/generated/`，请求会复用 `src/util/request.ts` 中的 Axios 实例。
