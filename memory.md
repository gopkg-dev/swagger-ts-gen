# Project Memory

- Requirement: build a Go CLI to generate TypeScript API files from Swagger/OpenAPI doc.
- Input source: local ./doc.json is present; production use requires fetching from URL.
- Current spec: doc.json is Swagger 2.0 with 19 tags and 78 definitions.
- OperationId coverage: 99 operations missing operationId; naming rule needed.
- Responses commonly use allOf with fiberx.Response and override data schema.
- Some operations use formData file parameters; upload handling rule needed.
- Grouping rule: group by path segment, skipping /api/v{n} prefix (e.g., /api/v1/users -> users).
- Naming rule: prefer operationId, fallback to METHOD + path (camelCase).
- Upload handling: detect formData and build FormData payload for request.
- Model output: only generate a single model/index.ts bundle (no per-type files), even if >500 lines (user confirmed).
- Pagination rule: query params with current + pageSize extend PageParam and exclude these fields from QueryParam body.
- Pagination return: when page query and data is array, return PageResult<T> and use ApiResult<PageResult<T>>.
- QueryParam naming: use operationId prefix stripping (query*), otherwise fall back to group name for QueryParam to match backend naming (e.g., RolesQueryParam).
- QueryParam field comments now mapped from Swagger param description when building query schema.
- FormData generation now guards against undefined body and preserves indentation in generated TS.
- Added verbose logging flag (-v/--verbose) to control runtime log output in CLI.
- Verbose logging now prints per-operation details (method/path/query/body/return).
- CLI framework switched to cobra (professional command line package).
