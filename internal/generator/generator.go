package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Options struct {
	OutputDir string
	Logf      func(string, ...any)
}

type Report struct {
	Groups     int
	Operations int
	Types      int
}

type Generator struct {
	spec      *openapi3.T
	outputDir string
	logf      func(string, ...any)
}

func New(spec *openapi3.T, opts Options) *Generator {
	output := strings.TrimSpace(opts.OutputDir)
	if output == "" {
		output = "api"
	}
	return &Generator{spec: spec, outputDir: output, logf: opts.Logf}
}

func (g *Generator) Generate() (*Report, error) {
	if g.spec == nil {
		return nil, fmt.Errorf("spec is nil")
	}

	ops, err := ExtractOperations(g.spec)
	if err != nil {
		return nil, err
	}

	groups := map[string][]RawOperation{}
	for _, op := range ops {
		groups[op.Group] = append(groups[op.Group], op)
	}

	groupNames := make([]string, 0, len(groups))
	for name := range groups {
		groupNames = append(groupNames, name)
	}
	sort.Strings(groupNames)

	if err := os.MkdirAll(g.outputDir, 0o755); err != nil {
		return nil, fmt.Errorf("create output dir failed: %w", err)
	}
	if err := os.WriteFile(filepath.Join(g.outputDir, "index.ts"), []byte(renderRootIndexFile()), 0o644); err != nil {
		return nil, fmt.Errorf("write root index failed: %w", err)
	}

	report := &Report{}

	for _, groupName := range groupNames {
		rawOps := groups[groupName]
		typedOps, apiImports, usesPageResult, registry, err := g.buildGroupOperations(rawOps)
		if err != nil {
			return nil, err
		}

		groupDir := filepath.Join(g.outputDir, groupName)
		modelDir := filepath.Join(groupDir, "model")

		if err := os.MkdirAll(modelDir, 0o755); err != nil {
			return nil, fmt.Errorf("create model dir failed: %w", err)
		}

		expandRegistryReferences(registry)
		typeDefs := registry.Types()
		report.Types += len(typeDefs)

		bundleContent, bundleLines := renderModelBundle(typeDefs, registry)
		if bundleLines > 0 {
			if err := os.WriteFile(filepath.Join(modelDir, "index.ts"), []byte(bundleContent), 0o644); err != nil {
				return nil, fmt.Errorf("write model bundle failed: %w", err)
			}
		}

		apiFiles := SplitAndRenderAPI(typedOps, apiImports, usesPageResult)
		for idx, content := range apiFiles {
			name := "index.ts"
			if len(apiFiles) > 1 {
				name = fmt.Sprintf("api_%d.ts", idx+1)
			}
			if err := os.WriteFile(filepath.Join(groupDir, name), []byte(content), 0o644); err != nil {
				return nil, fmt.Errorf("write api file failed: %w", err)
			}
		}

		if len(apiFiles) > 1 {
			indexContent := renderAPIIndex(len(apiFiles))
			if err := os.WriteFile(filepath.Join(groupDir, "index.ts"), []byte(indexContent), 0o644); err != nil {
				return nil, fmt.Errorf("write api index failed: %w", err)
			}
		}

		report.Groups++
		report.Operations += len(rawOps)
	}

	return report, nil
}

func renderTypeFile(content string, deps []string) string {
	if len(deps) == 0 {
		return content
	}
	var b strings.Builder
	b.WriteString("import type { ")
	b.WriteString(strings.Join(deps, ", "))
	b.WriteString(" } from './index';\n\n")
	b.WriteString(content)
	return b.String()
}

func renderAPIIndex(parts int) string {
	var b strings.Builder
	for i := 1; i <= parts; i++ {
		b.WriteString(fmt.Sprintf("export * from './api_%d';\n", i))
	}
	return b.String()
}

func (g *Generator) buildGroupOperations(rawOps []RawOperation) ([]Operation, []string, bool, *TypeRegistry, error) {
	registry := NewTypeRegistry(g.spec)
	usedTypes := map[string]struct{}{}
	usesPageResult := false

	ops := make([]Operation, 0, len(rawOps))
	for _, raw := range rawOps {
		op := Operation{
			Name:    ensureUniqueOperationName(raw.Name, ops),
			Summary: raw.Summary,
			Method:  raw.Method,
			Path:    raw.Path,
			Group:   raw.Group,
		}

		op.PathParams = buildPathParams(raw.PathParams, registry)
		for _, param := range op.PathParams {
			if _, ok := registry.types[param.Type]; ok {
				usedTypes[param.Type] = struct{}{}
			}
		}

		isPageQuery := hasPageParams(raw.QueryParams)
		if len(raw.QueryParams) > 0 {
			querySchema := buildQuerySchema(raw.QueryParams, isPageQuery)
			var typeName string
			queryParamName := buildQueryParamTypeName(op.Name, op.Group)
			if isPageQuery {
				typeName = registry.RegisterInlineWithExtends(queryParamName, querySchema, "", []string{"PageParam"})
			} else {
				typeName = registry.RegisterInline(queryParamName, querySchema, "")
			}
			op.Query = &QueryInfo{TypeName: typeName, Optional: !hasRequiredParams(raw.QueryParams, isPageQuery)}
			usedTypes[typeName] = struct{}{}
		}

		if raw.Body != nil {
			typeName := ""
			if raw.Body.Schema != nil && raw.Body.Schema.Ref != "" {
				refName, err := registry.RegisterRef(raw.Body.Schema.Ref)
				if err != nil {
					return nil, nil, false, nil, err
				}
				typeName = refName
			} else {
				desc := ""
				if raw.Body.IsFormData {
					desc = "FormData"
				}
				typeName = registry.RegisterInline(op.Name+"Body", raw.Body.Schema, desc)
			}
			op.Body = &BodyInfo{TypeName: typeName, Optional: !raw.Body.Required, IsForm: raw.Body.IsFormData}
			usedTypes[typeName] = struct{}{}
		}

		returnInfo, returnTypes := resolveReturnType(op.Name, raw.Response, registry, isPageQuery)
		op.Return = returnInfo
		if returnInfo.UsesPageResult {
			usesPageResult = true
		}
		for _, name := range returnTypes {
			usedTypes[name] = struct{}{}
		}

		op.ErrorText = buildErrorText(op.Summary)

		if g.logf != nil {
			g.logf(
				"operation name=%s method=%s path=%s group=%s page=%t pathParams=%s query=%s queryFields=%s body=%s return=%s",
				op.Name,
				strings.ToUpper(op.Method),
				op.Path,
				op.Group,
				isPageQuery,
				formatRawParamNames(raw.PathParams, false),
				formatQueryLog(op.Query, isPageQuery),
				formatRawParamNames(raw.QueryParams, isPageQuery),
				formatBodyLog(op.Body),
				op.Return.Type,
			)
		}

		ops = append(ops, op)
	}

	apiImports := make([]string, 0, len(usedTypes))
	for name := range usedTypes {
		apiImports = append(apiImports, name)
	}
	sort.Strings(apiImports)

	return ops, apiImports, usesPageResult, registry, nil
}

func ensureUniqueOperationName(name string, ops []Operation) string {
	used := map[string]struct{}{}
	for _, op := range ops {
		used[op.Name] = struct{}{}
	}
	if _, ok := used[name]; !ok {
		return name
	}
	idx := 2
	candidate := fmt.Sprintf("%s%d", name, idx)
	for {
		if _, ok := used[candidate]; !ok {
			return candidate
		}
		idx++
		candidate = fmt.Sprintf("%s%d", name, idx)
	}
}

func buildQuerySchema(params []RawParam, excludePage bool) *openapi3.SchemaRef {
	schema := &openapi3.Schema{Type: typesOf("object"), Properties: map[string]*openapi3.SchemaRef{}}
	for _, param := range params {
		if excludePage && isPageParamName(param.Name) {
			continue
		}
		propSchema := schemaOrAny(param.Schema)
		if param.Description != "" && propSchema.Ref == "" {
			if propSchema.Value == nil {
				propSchema.Value = &openapi3.Schema{}
			}
			propSchema.Value.Description = param.Description
		}
		schema.Properties[param.Name] = propSchema
		if param.Required {
			schema.Required = append(schema.Required, param.Name)
		}
	}
	return &openapi3.SchemaRef{Value: schema}
}

func hasRequiredParams(params []RawParam, excludePage bool) bool {
	for _, param := range params {
		if excludePage && isPageParamName(param.Name) {
			continue
		}
		if param.Required {
			return true
		}
	}
	return false
}

func hasPageParams(params []RawParam) bool {
	hasCurrent := false
	hasPageSize := false
	for _, param := range params {
		switch param.Name {
		case "current":
			hasCurrent = true
		case "pageSize":
			hasPageSize = true
		}
	}
	return hasCurrent && hasPageSize
}

func isPageParamName(name string) bool {
	return name == "current" || name == "pageSize"
}

const queryPrefix = "query"

func buildQueryParamTypeName(operationName string, groupName string) string {
	trimmed := strings.TrimSpace(operationName)
	if trimmed == "" {
		return "QueryParam"
	}
	if strings.EqualFold(trimmed, queryPrefix) {
		return "QueryParam"
	}
	lower := strings.ToLower(trimmed)
	if strings.HasPrefix(lower, queryPrefix) && len(trimmed) > len(queryPrefix) {
		base := trimmed[len(queryPrefix):]
		return base + "QueryParam"
	}
	group := strings.TrimSpace(groupName)
	if group != "" {
		return group + "QueryParam"
	}
	return trimmed + "QueryParam"
}

func buildPathParams(params []RawParam, registry *TypeRegistry) []Param {
	result := make([]Param, 0, len(params))
	usedNames := map[string]struct{}{}
	for _, param := range params {
		typeName := registry.SchemaToType(param.Schema, nil)
		varName := sanitizeIdentifier(param.Name)
		if varName == "data" || varName == "params" {
			varName = varName + "Param"
		}
		if _, ok := usedNames[varName]; ok {
			varName = fmt.Sprintf("%sParam", varName)
		}
		usedNames[varName] = struct{}{}

		result = append(result, Param{
			Name:        param.Name,
			VarName:     varName,
			Type:        typeName,
			Description: param.Description,
			Required:    param.Required,
		})
	}

	return result
}

func resolveReturnType(opName string, schemaRef *openapi3.SchemaRef, registry *TypeRegistry, isPageQuery bool) (ReturnInfo, []string) {
	dataSchema := extractDataSchema(schemaRef, registry)
	if dataSchema == nil || isEmptySchema(dataSchema) {
		return ReturnInfo{Type: "void", IsVoid: true}, nil
	}

	schema := dataSchema
	if schema.Ref != "" {
		if isPageQuery {
			resolved := derefSchemaRef(schema, registry)
			if resolved != nil && resolved.Value != nil && resolved.Value.Type != nil && resolved.Value.Type.Is("array") {
				itemType := registry.SchemaToType(resolved.Value.Items, nil)
				used := collectTypeNamesFromSchema(resolved.Value.Items, registry)
				return ReturnInfo{Type: "PageResult<" + itemType + ">", UsesPageResult: true}, used
			}
		}
		name, err := registry.RegisterRef(schema.Ref)
		if err != nil {
			return ReturnInfo{Type: "any", IsVoid: false}, nil
		}
		return ReturnInfo{Type: name}, []string{name}
	}

	if schema.Value == nil {
		return ReturnInfo{Type: "any"}, nil
	}

	if listItems := extractPageListItems(schema.Value, registry); listItems != nil {
		itemType := registry.SchemaToType(listItems, nil)
		used := collectTypeNamesFromSchema(listItems, registry)
		return ReturnInfo{Type: "PageResult<" + itemType + ">", UsesPageResult: true}, used
	}

	if schema.Value.Type != nil && schema.Value.Type.Is("array") {
		itemType := registry.SchemaToType(schema.Value.Items, nil)
		used := collectTypeNamesFromSchema(schema.Value.Items, registry)
		if isPageQuery {
			return ReturnInfo{Type: "PageResult<" + itemType + ">", UsesPageResult: true}, used
		}
		return ReturnInfo{Type: itemType + "[]"}, used
	}

	inlineName := registry.RegisterInline(opName+"Result", schema, "")
	return ReturnInfo{Type: inlineName}, []string{inlineName}
}

func extractPageListItems(schema *openapi3.Schema, registry *TypeRegistry) *openapi3.SchemaRef {
	if schema == nil || schema.Type == nil || !schema.Type.Is("object") {
		return nil
	}
	if schema.Properties == nil {
		return nil
	}
	listSchema := schema.Properties["list"]
	countSchema := schema.Properties["count"]
	if listSchema == nil || countSchema == nil {
		return nil
	}
	listSchema = derefSchemaRef(listSchema, registry)
	if listSchema == nil || listSchema.Value == nil {
		return nil
	}
	if listSchema.Value.Type == nil || !listSchema.Value.Type.Is("array") {
		return nil
	}
	return listSchema.Value.Items
}

func collectTypeNamesFromSchema(schemaRef *openapi3.SchemaRef, registry *TypeRegistry) []string {
	if schemaRef == nil {
		return nil
	}
	if schemaRef.Ref != "" {
		name, err := registry.RegisterRef(schemaRef.Ref)
		if err != nil {
			return nil
		}
		return []string{name}
	}

	schema := schemaRef.Value
	if schema == nil {
		return nil
	}

	var names []string
	if schema.Items != nil {
		names = append(names, collectTypeNamesFromSchema(schema.Items, registry)...)
	}

	for _, prop := range schema.Properties {
		names = append(names, collectTypeNamesFromSchema(prop, registry)...)
	}

	return uniqueStrings(names)
}

func derefSchemaRef(schemaRef *openapi3.SchemaRef, registry *TypeRegistry) *openapi3.SchemaRef {
	if schemaRef == nil || registry == nil {
		return schemaRef
	}
	if schemaRef.Ref == "" {
		return schemaRef
	}
	resolved := registry.resolveRefSchema(schemaRef.Ref)
	if resolved != nil {
		return resolved
	}
	return schemaRef
}

func uniqueStrings(items []string) []string {
	seen := map[string]struct{}{}
	var result []string
	for _, item := range items {
		if item == "" {
			continue
		}
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	sort.Strings(result)
	return result
}

func extractDataSchema(schemaRef *openapi3.SchemaRef, registry *TypeRegistry) *openapi3.SchemaRef {
	if schemaRef == nil {
		return nil
	}
	if schemaRef.Ref != "" {
		refSchema := registry.resolveRefSchema(schemaRef.Ref)
		if refSchema != nil {
			schemaRef = refSchema
		}
	}
	if schemaRef == nil || schemaRef.Value == nil {
		return nil
	}

	schema := schemaRef.Value
	if len(schema.AllOf) > 0 {
		for _, item := range schema.AllOf {
			if item == nil || item.Value == nil {
				continue
			}
			if dataSchema, ok := item.Value.Properties["data"]; ok {
				if !isEmptySchema(dataSchema) {
					return dataSchema
				}
			}
		}
	}

	if dataSchema, ok := schema.Properties["data"]; ok {
		if !isEmptySchema(dataSchema) {
			return dataSchema
		}
	}

	return nil
}

func isEmptySchema(schemaRef *openapi3.SchemaRef) bool {
	if schemaRef == nil {
		return true
	}
	if schemaRef.Ref != "" {
		return false
	}
	schema := schemaRef.Value
	if schema == nil {
		return true
	}
	if schema.Type != nil && len(*schema.Type) > 0 {
		return false
	}
	if len(schema.Properties) > 0 || schema.Items != nil || len(schema.AllOf) > 0 || len(schema.OneOf) > 0 || len(schema.AnyOf) > 0 || len(schema.Enum) > 0 {
		return false
	}
	return true
}

func buildErrorText(summary string) string {
	trimmed := strings.TrimSpace(summary)
	if trimmed == "" {
		return "请求失败"
	}
	if strings.HasSuffix(trimmed, "失败") {
		return trimmed
	}
	return trimmed + "失败"
}

func formatRawParamNames(params []RawParam, excludePage bool) string {
	if len(params) == 0 {
		return "-"
	}
	names := make([]string, 0, len(params))
	for _, param := range params {
		if excludePage && isPageParamName(param.Name) {
			continue
		}
		if param.Name == "" {
			continue
		}
		names = append(names, param.Name)
	}
	if len(names) == 0 {
		return "-"
	}
	return strings.Join(names, ",")
}

func formatQueryLog(query *QueryInfo, isPageQuery bool) string {
	if query == nil {
		return "-"
	}
	suffix := ""
	if query.Optional {
		suffix = "?"
	}
	pageSuffix := ""
	if isPageQuery {
		pageSuffix = "+PageParam"
	}
	return query.TypeName + suffix + pageSuffix
}

func formatBodyLog(body *BodyInfo) string {
	if body == nil {
		return "-"
	}
	suffix := ""
	if body.Optional {
		suffix = "?"
	}
	formSuffix := ""
	if body.IsForm {
		formSuffix = "(form)"
	}
	return body.TypeName + suffix + formSuffix
}

func SplitAndRenderAPI(ops []Operation, modelImports []string, usesPageResult bool) []string {
	if len(ops) == 0 {
		return nil
	}
	maxLines := 500
	opStrings := make([]string, 0, len(ops))
	opLines := make([]int, 0, len(ops))

	for _, op := range ops {
		content := RenderOperation(op)
		opStrings = append(opStrings, content)
		opLines = append(opLines, countLines(content))
	}

	header := RenderAPIFile([]Operation{}, modelImports, usesPageResult)
	headerLines := countLines(header)

	var files []string
	var buffer []string
	lineCount := headerLines

	for idx, content := range opStrings {
		lines := opLines[idx]
		if lineCount+lines > maxLines && len(buffer) > 0 {
			files = append(files, renderAPIFromChunks(buffer, header, headerLines))
			buffer = nil
			lineCount = headerLines
		}
		buffer = append(buffer, content)
		lineCount += lines
	}

	if len(buffer) > 0 {
		files = append(files, renderAPIFromChunks(buffer, header, headerLines))
	}

	return files
}

func renderAPIFromChunks(chunks []string, header string, headerLines int) string {
	if len(chunks) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(header)
	if headerLines > 0 && !strings.HasSuffix(header, "\n\n") {
		b.WriteString("\n")
	}
	for idx, chunk := range chunks {
		if idx > 0 {
			b.WriteString("\n")
		}
		b.WriteString(chunk)
		b.WriteString("\n")
	}
	return b.String()
}

func countLines(content string) int {
	if content == "" {
		return 0
	}
	return strings.Count(content, "\n") + 1
}

func typesOf(values ...string) *openapi3.Types {
	if len(values) == 0 {
		return nil
	}
	types := openapi3.Types(values)
	return &types
}

func extractTypeNames(defs []*TypeDef) []string {
	names := make([]string, 0, len(defs))
	for _, def := range defs {
		if def == nil || def.Name == "" {
			continue
		}
		names = append(names, def.Name)
	}
	sort.Strings(names)
	return names
}

func renderModelBundle(defs []*TypeDef, registry *TypeRegistry) (string, int) {
	if len(defs) == 0 {
		return "", 0
	}
	var b strings.Builder
	lineCount := 0
	if needsPageParamImport(defs) {
		b.WriteString("import type { PageParam } from '@/api';\n\n")
		lineCount += 2
	}
	for idx, def := range defs {
		content, _ := RenderType(def, registry)
		if content == "" {
			continue
		}
		if idx > 0 {
			b.WriteString("\n")
			lineCount++
		}
		b.WriteString(content)
		lineCount += countLines(content)
	}
	return b.String(), lineCount
}

func expandRegistryReferences(registry *TypeRegistry) {
	if registry == nil {
		return
	}
	visitedRefs := map[string]struct{}{}
	visitedSchemas := map[*openapi3.Schema]struct{}{}
	for {
		before := len(registry.typeOrder)
		defs := registry.Types()
		for _, def := range defs {
			walkSchemaRefs(def.Schema, registry, visitedRefs, visitedSchemas)
		}
		after := len(registry.typeOrder)
		if after == before {
			return
		}
	}
}

func needsPageParamImport(defs []*TypeDef) bool {
	for _, def := range defs {
		if def == nil {
			continue
		}
		for _, ext := range def.Extends {
			if ext == "PageParam" {
				return true
			}
		}
	}
	return false
}

func walkSchemaRefs(schemaRef *openapi3.SchemaRef, registry *TypeRegistry, visitedRefs map[string]struct{}, visitedSchemas map[*openapi3.Schema]struct{}) {
	if schemaRef == nil || registry == nil {
		return
	}
	if schemaRef.Ref != "" {
		if _, ok := visitedRefs[schemaRef.Ref]; ok {
			return
		}
		visitedRefs[schemaRef.Ref] = struct{}{}
		_, _ = registry.RegisterRef(schemaRef.Ref)
		if resolved := registry.resolveRefSchema(schemaRef.Ref); resolved != nil {
			walkSchemaRefs(resolved, registry, visitedRefs, visitedSchemas)
		}
		return
	}
	schema := schemaRef.Value
	if schema == nil {
		return
	}
	if _, ok := visitedSchemas[schema]; ok {
		return
	}
	visitedSchemas[schema] = struct{}{}
	for _, ref := range schema.AllOf {
		walkSchemaRefs(ref, registry, visitedRefs, visitedSchemas)
	}
	for _, ref := range schema.OneOf {
		walkSchemaRefs(ref, registry, visitedRefs, visitedSchemas)
	}
	for _, ref := range schema.AnyOf {
		walkSchemaRefs(ref, registry, visitedRefs, visitedSchemas)
	}
	if schema.Not != nil {
		walkSchemaRefs(schema.Not, registry, visitedRefs, visitedSchemas)
	}
	if schema.Items != nil {
		walkSchemaRefs(schema.Items, registry, visitedRefs, visitedSchemas)
	}
	for _, prop := range schema.Properties {
		walkSchemaRefs(prop, registry, visitedRefs, visitedSchemas)
	}
	if schema.AdditionalProperties.Schema != nil {
		walkSchemaRefs(schema.AdditionalProperties.Schema, registry, visitedRefs, visitedSchemas)
	}
}
