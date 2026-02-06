package generator

import (
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func RenderType(def *TypeDef, registry *TypeRegistry) (string, []string) {
	deps := map[string]struct{}{}

	schema := def.Schema
	if schema == nil || schema.Value == nil && schema.Ref == "" {
		content := fmt.Sprintf("export type %s = any;\n", def.Name)
		return content, nil
	}

	if schema.Ref != "" {
		_, _ = registry.RegisterRef(schema.Ref)
	}

	resolved := schema
	if schema.Ref != "" {
		if refSchema := registry.resolveRefSchema(schema.Ref); refSchema != nil {
			resolved = refSchema
		}
	}

	content := renderTypeDefinition(def, resolved, registry, deps)

	depList := make([]string, 0, len(deps))
	for name := range deps {
		if name == def.Name {
			continue
		}
		depList = append(depList, name)
	}
	sort.Strings(depList)

	return content, depList
}

func renderTypeDefinition(def *TypeDef, schemaRef *openapi3.SchemaRef, registry *TypeRegistry, deps map[string]struct{}) string {
	if schemaRef == nil || schemaRef.Value == nil {
		return fmt.Sprintf("export type %s = any;\n", def.Name)
	}

	schema := schemaRef.Value
	description := strings.TrimSpace(def.Description)
	if description == "" {
		description = strings.TrimSpace(schema.Description)
	}

	isObject := (schema.Type != nil && schema.Type.Is("object")) || len(schema.Properties) > 0 || schema.AdditionalProperties.Schema != nil || schema.AdditionalProperties.Has != nil
	if len(schema.Enum) > 0 || len(schema.OneOf) > 0 || len(schema.AnyOf) > 0 || len(schema.AllOf) > 0 || !isObject {
		typeExpr := registry.schemaValueToType(schema, deps)
		if schema.Nullable {
			typeExpr = typeExpr + " | null"
		}
		return formatTypeAlias(def.Name, typeExpr, description)
	}

	return formatInterface(def.Name, schema, registry, deps, description, def.Extends)
}

func formatTypeAlias(name string, expr string, description string) string {
	var b strings.Builder
	if description != "" {
		b.WriteString("/** " + description + " */\n")
	}
	b.WriteString("export type " + name + " = " + expr + ";\n")
	return b.String()
}

func formatInterface(name string, schema *openapi3.Schema, registry *TypeRegistry, deps map[string]struct{}, description string, extends []string) string {
	required := map[string]struct{}{}
	for _, key := range schema.Required {
		required[key] = struct{}{}
	}

	var keys []string
	for key := range schema.Properties {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var b strings.Builder
	if description != "" {
		b.WriteString("/** " + description + " */\n")
	}
	extendClause := ""
	if len(extends) > 0 {
		extendClause = " extends " + strings.Join(extends, ", ")
	}
	b.WriteString("export interface " + name + extendClause + " {\n")
	for _, key := range keys {
		propSchema := schema.Properties[key]
		if propSchema == nil {
			continue
		}
		optional := "?"
		if _, ok := required[key]; ok {
			optional = ""
		}
		propName := key
		if !isValidIdentifier(key) {
			propName = fmt.Sprintf("'%s'", escapeSingleQuotes(key))
		}
		propDesc := ""
		if propSchema.Value != nil {
			propDesc = strings.TrimSpace(propSchema.Value.Description)
		}
		if propDesc != "" {
			b.WriteString("  /** " + propDesc + " */\n")
		}
		propType := registry.SchemaToType(propSchema, deps)
		b.WriteString("  " + propName + optional + ": " + propType + ";\n")
	}

	if len(schema.Properties) == 0 && schema.AdditionalProperties.Schema != nil {
		valueType := registry.SchemaToType(schema.AdditionalProperties.Schema, deps)
		b.WriteString("  [key: string]: " + valueType + ";\n")
	} else if len(schema.Properties) == 0 && schema.AdditionalProperties.Has != nil && *schema.AdditionalProperties.Has {
		b.WriteString("  [key: string]: any;\n")
	}

	b.WriteString("}\n")
	return b.String()
}

func RenderModelIndex(typeNames []string) string {
	if len(typeNames) == 0 {
		return ""
	}
	sort.Strings(typeNames)
	var b strings.Builder
	for _, name := range typeNames {
		b.WriteString("export type { " + name + " } from './" + name + "';\n")
	}
	return b.String()
}

func RenderAPIFile(ops []Operation, modelImports []string, usesPageResult bool) string {
	var b strings.Builder
	b.WriteString("import request from '@/utils/request';\n")

	apiImports := []string{"ApiResult"}
	if usesPageResult {
		apiImports = append(apiImports, "PageResult")
	}
	b.WriteString("import type { " + strings.Join(apiImports, ", ") + " } from '@/api';\n")

	if len(modelImports) > 0 {
		sort.Strings(modelImports)
		b.WriteString("import type {\n")
		for _, name := range modelImports {
			b.WriteString("  " + name + ",\n")
		}
		b.WriteString("} from './model';\n")
	}

	b.WriteString("\n")

	for idx, op := range ops {
		if idx > 0 {
			b.WriteString("\n")
		}
		b.WriteString(RenderOperation(op))
		b.WriteString("\n")
	}

	return b.String()
}

func RenderOperation(op Operation) string {
	var b strings.Builder
	summary := strings.TrimSpace(op.Summary)
	if summary == "" {
		summary = op.Name
	}

	b.WriteString("/**\n")
	b.WriteString(" * " + summary + "\n")
	for _, param := range op.PathParams {
		if param.Description == "" {
			b.WriteString(" * @param " + param.VarName + " - 路径参数\n")
		} else {
			b.WriteString(" * @param " + param.VarName + " - " + param.Description + "\n")
		}
	}
	if op.Body != nil {
		b.WriteString(" * @param data - 请求数据\n")
	}
	if op.Query != nil {
		b.WriteString(" * @param params - 查询参数\n")
	}
	b.WriteString(" * @returns Promise<" + op.Return.Type + ">\n")
	b.WriteString(" */\n")

	args := renderOperationArgs(op)
	b.WriteString("export async function " + op.Name + "(" + args + ") {\n")

	url := renderPathTemplate(op)
	if op.Body != nil && op.Body.IsForm {
		b.WriteString("  const formData = new FormData();\n")
		b.WriteString("  if (data) {\n")
		b.WriteString("    for (const [key, value] of Object.entries(data)) {\n")
		b.WriteString("      if (value === undefined || value === null) {\n")
		b.WriteString("        continue;\n")
		b.WriteString("      }\n")
		b.WriteString("      if (Array.isArray(value)) {\n")
		b.WriteString("        for (const item of value) {\n")
		b.WriteString("          if (item !== undefined && item !== null) {\n")
		b.WriteString("            formData.append(key, item as any);\n")
		b.WriteString("          }\n")
		b.WriteString("        }\n")
		b.WriteString("      } else {\n")
		b.WriteString("        formData.append(key, value as any);\n")
		b.WriteString("      }\n")
		b.WriteString("    }\n")
		b.WriteString("  }\n")
	}

	reqLine := renderRequest(op, url)
	b.WriteString(reqLine)

	if op.Return.IsVoid {
		b.WriteString("  if (res.data.success) {\n")
		b.WriteString("    return;\n")
		b.WriteString("  }\n")
		b.WriteString("  return Promise.reject(new Error(res.data.error?.message ?? '")
		b.WriteString(escapeSingleQuotes(op.ErrorText))
		b.WriteString("'));\n")
	} else {
		b.WriteString("  if (res.data.success && res.data.data !== undefined) {\n")
		b.WriteString("    return res.data.data;\n")
		b.WriteString("  }\n")
		b.WriteString("  return Promise.reject(new Error(res.data.error?.message ?? '")
		b.WriteString(escapeSingleQuotes(op.ErrorText))
		b.WriteString("'));\n")
	}

	b.WriteString("}\n")

	return b.String()
}

func renderOperationArgs(op Operation) string {
	var args []string
	for _, param := range op.PathParams {
		name := sanitizeIdentifier(param.VarName)
		arg := name + ": " + param.Type
		if !param.Required {
			arg = name + "?: " + param.Type
		}
		args = append(args, arg)
	}
	if op.Body != nil {
		if op.Body.Optional {
			args = append(args, "data?: "+op.Body.TypeName)
		} else {
			args = append(args, "data: "+op.Body.TypeName)
		}
	}
	if op.Query != nil {
		if op.Query.Optional {
			args = append(args, "params?: "+op.Query.TypeName)
		} else {
			args = append(args, "params: "+op.Query.TypeName)
		}
	}

	return strings.Join(args, ", ")
}

func renderPathTemplate(op Operation) string {
	if len(op.PathParams) == 0 {
		return "'" + escapeSingleQuotes(op.Path) + "'"
	}
	path := op.Path
	for _, param := range op.PathParams {
		placeholder := "{" + param.Name + "}"
		path = strings.ReplaceAll(path, placeholder, "${"+param.VarName+"}")
	}
	return "`" + path + "`"
}

func renderRequest(op Operation, url string) string {
	method := strings.ToLower(op.Method)
	returnType := op.Return.Type

	if op.Body != nil && op.Body.IsForm {
		config := buildConfigObject(op, false, true)
		if config == "" {
			return fmt.Sprintf("  const res = await request.%s<ApiResult<%s>>(%s, formData);\n", method, returnType, url)
		}
		return fmt.Sprintf("  const res = await request.%s<ApiResult<%s>>(%s, formData, %s);\n", method, returnType, url, config)
	}

	if op.Body != nil {
		if method == "delete" {
			config := buildConfigObject(op, true, false)
			if config == "" {
				return fmt.Sprintf("  const res = await request.delete<ApiResult<%s>>(%s);\n", returnType, url)
			}
			return fmt.Sprintf("  const res = await request.delete<ApiResult<%s>>(%s, %s);\n", returnType, url, config)
		}

		config := buildConfigObject(op, false, false)
		if config == "" {
			return fmt.Sprintf("  const res = await request.%s<ApiResult<%s>>(%s, data);\n", method, returnType, url)
		}
		return fmt.Sprintf("  const res = await request.%s<ApiResult<%s>>(%s, data, %s);\n", method, returnType, url, config)
	}

	if op.Query != nil {
		config := buildConfigObject(op, false, false)
		return fmt.Sprintf("  const res = await request.%s<ApiResult<%s>>(%s, %s);\n", method, returnType, url, config)
	}

	return fmt.Sprintf("  const res = await request.%s<ApiResult<%s>>(%s);\n", method, returnType, url)
}

func buildConfigObject(op Operation, includeData bool, includeFormHeader bool) string {
	var entries []string
	if includeData {
		entries = append(entries, "data")
	}
	if op.Query != nil {
		entries = append(entries, "params")
	}
	if includeFormHeader {
		entries = append(entries, "headers: { 'Content-Type': 'multipart/form-data' }")
	}
	if len(entries) == 0 {
		return ""
	}
	return "{ " + strings.Join(entries, ", ") + " }"
}

func escapeSingleQuotes(value string) string {
	if value == "" {
		return value
	}
	escaped := strings.ReplaceAll(value, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "'", "\\'")
	escaped = strings.ReplaceAll(escaped, "\n", " ")
	return escaped
}
