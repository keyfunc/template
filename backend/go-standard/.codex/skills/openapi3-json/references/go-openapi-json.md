# Go OpenAPI JSON Reference

Use this reference when generating OpenAPI 3 JSON from Go backend code.

## Route Discovery

Search with `rg` first:

```bash
rg -n 'HandleFunc|Handle\(|http\.Method|PathValue|URL\.Query|json\.NewDecoder|Bind|ShouldBind|Param\(|Query\(' --glob '*.go'
```

Common patterns:

- Standard library Go 1.22 `ServeMux`: `mux.HandleFunc("GET /api/v1/todo/{id}", handler.Update)`.
- Standard `ServeMux` before method patterns: inspect handler branches for `r.Method`.
- `r.PathValue("id")`: path parameter named `id`.
- `r.URL.Query().Get("page")`: query parameter named `page`.
- `json.NewDecoder(r.Body).Decode(&req)`: JSON request body schema comes from `req`.

## Go Type Mapping

Map Go types to OpenAPI schema types:

| Go type | OpenAPI schema |
| --- | --- |
| `string` | `{ "type": "string" }` |
| `int`, `int32`, `uint`, `uint32` | `{ "type": "integer", "format": "int32" }` |
| `int64`, `uint64` | `{ "type": "integer", "format": "int64" }` |
| `float32` | `{ "type": "number", "format": "float" }` |
| `float64` | `{ "type": "number", "format": "double" }` |
| `bool` | `{ "type": "boolean" }` |
| `time.Time` | `{ "type": "string", "format": "date-time" }` |
| `[]T` | `{ "type": "array", "items": schema-for-T }` |
| `*T` | Same as `T`, with `nullable: true` for OpenAPI 3.0 |
| `map[string]T` | `{ "type": "object", "additionalProperties": schema-for-T }` |

Use JSON tag names exactly. Skip fields tagged `json:"-"`. For `omitempty`, do not mark the field required unless validation logic requires it.

## Required Fields

Mark a request field as required when one of these is true:

- Validation logic rejects the zero value, empty string, nil pointer, or missing value.
- A validation tag indicates required behavior.
- The handler cannot proceed without the field and returns an error.

Do not mark response fields required unless the response helper always emits them.

## Response Wrappers

Many Go APIs wrap all responses in a common envelope. Model the envelope once and reference concrete data schemas through composed or named schemas.

For a wrapper like:

```go
type Response[T any] struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    T      `json:"data"`
}
```

Create named schemas such as `TodoResponse`, `TodoListPageResponse`, and `EmptyResponse`, each with:

```json
{
  "type": "object",
  "required": ["code", "message", "data"],
  "properties": {
    "code": { "type": "integer", "format": "int32" },
    "message": { "type": "string" },
    "data": { "$ref": "#/components/schemas/Todo" }
  }
}
```

If failure helpers return HTTP 200 with application `code` set to an error value, document the HTTP status as `200` and describe the application-level error in the response description or an error envelope schema. Do not change it to HTTP 400 unless the code actually writes HTTP 400.

## Pagination

For helpers like:

```go
type PageData[T any] struct {
    List  []T   `json:"list"`
    Page  int   `json:"page"`
    Size  int   `json:"size"`
    Total int64 `json:"total"`
}
```

Use a named schema per resource, for example `TodoPageData`, and wrap it in the normal success envelope if the helper does that.

Query params commonly include:

- `page`: integer, minimum 1, default 1 when the code sets that default.
- `size`: integer, minimum 1, default 10 when the code sets that default.

## Operation Shape

Each operation should include:

- `tags`: resource/domain tag, for example `["todo"]`.
- `summary`: short human-readable action.
- `operationId`: stable lower-camel verb phrase.
- `parameters`: path and query params.
- `requestBody`: only for methods that read a JSON body.
- `responses`: at least the success response and any clearly documented error behavior.

Example path item:

```json
{
  "get": {
    "tags": ["todo"],
    "summary": "List todos",
    "operationId": "listTodos",
    "parameters": [
      {
        "name": "page",
        "in": "query",
        "required": false,
        "schema": { "type": "integer", "format": "int32", "minimum": 1, "default": 1 }
      }
    ],
    "responses": {
      "200": {
        "description": "Success response",
        "content": {
          "application/json": {
            "schema": { "$ref": "#/components/schemas/TodoPageResponse" }
          }
        }
      }
    }
  }
}
```

## JSON Document Shape

Use this top-level shape for a new file:

```json
{
  "openapi": "3.0.3",
  "info": {
    "title": "Project API",
    "version": "0.1.0"
  },
  "servers": [
    {
      "url": "http://localhost:8080"
    }
  ],
  "paths": {},
  "components": {
    "schemas": {}
  }
}
```

Infer `servers[0].url` from config only when obvious. Otherwise use `/` or a local development URL and mention the assumption.

## Validation Checklist

Before finishing:

- Run JSON syntax validation.
- Check every discovered route appears exactly once.
- Check every `$ref` target exists.
- Check path params in URLs have matching `in: path` parameters with `required: true`.
- Check request body schemas match DTO JSON tags and validation rules.
- Check response wrappers match actual response helpers.
