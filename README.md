# swagger-ts-gen

`swagger-ts-gen` 是一个 Go CLI，用于把 Swagger/OpenAPI 文档生成为前端可直接调用的 TypeScript API 函数与类型定义。

> ⚠️ 适用范围说明：该生成器当前仅面向 `go-fiber-admin` 项目约定设计，不保证可直接用于其他项目。

## 功能特性

- 支持输入来源：本地文件路径或 HTTP/HTTPS URL
- 支持文档格式：JSON / YAML
- 支持规范版本：OpenAPI 3、Swagger 2.0（自动转换为 OpenAPI 3 再处理）
- 按路径分组输出 API 文件，生成稳定、可复现的函数顺序
- 自动提取请求参数、查询参数、请求体、响应类型
- 自动处理分页场景（`current + pageSize`）并映射为 `PageResult<T>`
- 自动处理 `multipart/form-data` / `x-www-form-urlencoded` 请求并构造 `FormData`
- 自动处理引用类型与内联类型，按分组输出 `model/index.ts`

## 环境要求

- Go >= 1.25

## 快速开始

### 1) 安装（go install）

```bash
go install github.com/gopkg-dev/swagger-ts-gen/cmd/swagger-ts@latest
```

### 2) 运行生成器

```bash
swagger-ts -i ./doc.json -o ./api
```

示例输出：

```text
Source: ./doc.json
Spec: Swagger 2.0
Groups: 18, Operations: 99, Types: 110
```

### 3) 查看帮助

```bash
swagger-ts --help
```

## CLI 参数

- `-i, --input`：Swagger/OpenAPI 文档路径或 URL（必填）
- `-o, --output`：输出目录（默认 `api`）
- `-v, --verbose`：开启详细日志
- `--go-source`：Go 源码目录（用于 AST 可选性推断）
- `--go-source-include`：AST 扫描目录名（逗号分隔，默认 `schema,fiberx`）
- `--required-by-omitempty`：对象字段默认必填，仅 `omitempty` 字段输出可选（需配合 `--go-source`）

缺少 `--input` 时会以退出码 `2` 退出；其他错误为退出码 `1`。

## 输出结构

生成结果按分组落盘，典型结构如下：

```text
api/
  roles/
    index.ts
    model/
      index.ts
  users/
    index.ts
    model/
      index.ts
  ...
```

当某个分组的 API 文件超过 500 行时，会自动拆分为多个分片：

```text
api/<group>/
  api_1.ts
  api_2.ts
  ...
  index.ts   # 导出所有 api_n.ts
  model/
    index.ts
```

## 生成规则（核心）

### 1) 分组规则

- 路径形如 `/api/v1/users`：分组为 `users`（跳过 `/api/v{n}` 前缀）
- 其他路径：取第一个路径段作为分组名
- 分组名会标准化为 lowerCamel（例如 `sys-api -> sysApi`、`dict-items -> dictItems`）

### 2) API 函数命名

- 优先使用 `operationId`
- 若缺失，则使用 `method + path` 规则生成 lowerCamelCase 名称

### 3) 生成顺序（可复现）

- 先按路径字典序排序
- 同一路径内按固定 HTTP 方法顺序输出：
  `GET -> POST -> PUT -> PATCH -> DELETE -> HEAD -> OPTIONS`

### 4) 分页规则

- 查询参数同时包含 `current` 和 `pageSize` 时：
  - 查询参数类型扩展 `PageParam`
  - 返回数组会映射成 `PageResult<T>`

### 5) 请求体规则

- `multipart/form-data` 与 `application/x-www-form-urlencoded` 自动转 `FormData`
- 非表单请求按普通 JSON 体生成

### 6) 可选字段推断（可选能力）

- 默认行为：仅根据 OpenAPI `required` 数组决定 TS 字段是否可选。
- 开启 `--required-by-omitempty --go-source <dir>` 后：
  - 先用 Go AST 解析结构体 `json` tag
  - 扫描目录由 `--go-source-include` 决定（默认 `schema,fiberx`）
  - 当 schema 缺失 `required` 时，按“默认必填、仅 `omitempty` 可选”生成 TS 字段。
  - 若能匹配到 Go 结构体，会按 Go 字段声明顺序输出 TS 字段，前端阅读与后端定义保持一致。
  - 若存在同名结构体冲突（不同包同名），会基于字段重叠度选择最匹配 schema 的结构体进行覆盖。

## 生成代码依赖约定

生成的 TS 代码默认依赖以下项目约定：

- `@/utils/request`：统一请求实例
- `@/api`：导出 `ApiResult`（分页场景还需 `PageResult` 与 `PageParam`）

若你的工程别名或基础类型命名不同，请在接入前调整模板或统一适配层。

## 项目适配声明

- 当前实现绑定 `go-fiber-admin` 的接口返回结构与前端工程约定。
- 非 `go-fiber-admin` 项目请勿直接使用，需先改造模板与类型映射规则。

## 目录说明

- `cmd/swagger-ts`：CLI 入口
- `internal/loader`：文档读取与版本处理
- `internal/generator`：类型与 API 代码生成逻辑
