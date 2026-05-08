---
name: openapi3-json
description: Generate or update OpenAPI 3.x interface specifications as JSON from backend API code. Use when Codex needs to create docs/openapi.json or another OpenAPI JSON file, infer endpoints from Go net/http routes or router registrations, map handler request and response DTOs into schemas, document response envelopes, validate OpenAPI JSON, or maintain API documentation for project APIs.
---

# OpenAPI 3 JSON

## Overview

Use this skill to produce project-local OpenAPI 3 interface docs, with JSON as the primary output format. Prefer OpenAPI `3.0.3` unless the project already uses `3.1.x` or the user asks for it.

Default output path: `docs/openapi.json`, unless the project already has an OpenAPI file or the user names a different path.

## Workflow

1. Locate any existing API spec before creating a new one:
   - Search for `openapi.json`, `swagger.json`, `openapi.yaml`, `swagger.yaml`, `api-docs`, or `docs/`.
   - Preserve existing style, naming, tags, and schema structure when updating.

2. Inventory routes from source code:
   - For Go standard `http.ServeMux`, search `HandleFunc`, `Handle`, and Go 1.22 patterns such as `"GET /api/v1/todo/{id}"`.
   - For other routers, search route registration calls and grouped route prefixes.
   - Record method, path, handler function, module/domain, and route prefix.

3. Read the handler and related DTO/model files:
   - Identify path params from route templates or calls such as `r.PathValue("id")`.
   - Identify query params from `r.URL.Query()`, helper functions, DTOs, or validation logic.
   - Identify JSON request bodies from decoder/bind calls and request DTO structs.
   - Identify response bodies from response helpers, success paths, error paths, and pagination helpers.

4. Build or update a valid OpenAPI JSON document:
   - Include `openapi`, `info`, `servers`, `paths`, and `components.schemas`.
   - Use stable `operationId` values such as `listTodos`, `createTodo`, `updateTodo`, `deleteTodo`.
   - Use tags from the domain/package/resource name.
   - Put shared DTOs, models, response wrappers, and pagination envelopes in `components.schemas`.
   - Use `$ref` for repeated schemas; avoid duplicating object definitions inline.

5. Validate the JSON:
   - Run `python3 -m json.tool <file>` or `jq . <file>` to catch JSON syntax errors.
   - If the repo already has OpenAPI tooling, run the existing validation command.
   - Do not install generators or download validators unless the user approves or the repo already has that dependency.

## Output Rules

- Emit JSON, not YAML, unless the user explicitly requests YAML or an existing YAML spec must be updated.
- Do not invent undocumented endpoints, enum values, auth schemes, or business error codes.
- If a behavior is inferred from code rather than explicit docs, make the best conservative choice in the spec and mention the assumption in the final response.
- Treat application-level response wrappers separately from HTTP status codes. If the code returns HTTP 200 for failures with an application `code`, document that accurately.
- Include `required` only when validation or code clearly requires a field.
- Preserve JSON field names from `json` tags. If a Go field has no tag, infer the lower camel/snake name only when the surrounding code makes that convention clear.
- Prefer concise examples only when they can be derived from code, tests, fixtures, or config.

## Go Reference

For Go API projects, read `references/go-openapi-json.md` before writing or updating the spec. It contains route extraction hints, Go-to-OpenAPI type mapping, response wrapper patterns, and validation checks.

## Final Response

After generating or updating the spec, summarize:
- Output file path.
- Endpoint count and tags.
- Validation command and result.
- Any assumptions or ambiguous API behaviors.
