package generator

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type RawParam struct {
	Name        string
	In          string
	Description string
	Required    bool
	Schema      *openapi3.SchemaRef
}

type RawBody struct {
	Schema     *openapi3.SchemaRef
	Required   bool
	IsFormData bool
}

type RawOperation struct {
	Name        string
	Summary     string
	Method      string
	Path        string
	Group       string
	PathParams  []RawParam
	QueryParams []RawParam
	Body        *RawBody
	Response    *openapi3.SchemaRef
}

var pathParamRegexp = regexp.MustCompile(`\{([^}]+)\}`)

func ExtractOperations(doc *openapi3.T) ([]RawOperation, error) {
	if doc.Paths == nil {
		return nil, fmt.Errorf("spec has no paths")
	}

	pathItems := doc.Paths.Map()
	paths := make([]string, 0, len(pathItems))
	for path := range pathItems {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	var ops []RawOperation
	for _, path := range paths {
		item := pathItems[path]
		if item == nil {
			continue
		}
		for method, op := range operationsForPathItem(item) {
			if op == nil {
				continue
			}

			opName := sanitizeOperationID(op.OperationID)
			if opName == "" {
				opName = buildOperationName(method, path)
			}

			params := mergeParameters(doc, item.Parameters, op.Parameters)
			pathParams := extractPathParams(params, path)
			queryParams := extractQueryParams(params)

			body, err := extractRequestBody(op)
			if err != nil {
				return nil, err
			}

			responseSchema := extractResponseSchema(op)

			ops = append(ops, RawOperation{
				Name:        opName,
				Summary:     strings.TrimSpace(op.Summary),
				Method:      method,
				Path:        path,
				Group:       groupFromPath(path),
				PathParams:  pathParams,
				QueryParams: queryParams,
				Body:        body,
				Response:    responseSchema,
			})
		}
	}

	return ops, nil
}

func operationsForPathItem(item *openapi3.PathItem) map[string]*openapi3.Operation {
	return map[string]*openapi3.Operation{
		"get":     item.Get,
		"post":    item.Post,
		"put":     item.Put,
		"patch":   item.Patch,
		"delete":  item.Delete,
		"head":    item.Head,
		"options": item.Options,
	}
}

func mergeParameters(doc *openapi3.T, pathParams openapi3.Parameters, opParams openapi3.Parameters) map[string]*openapi3.Parameter {
	merged := map[string]*openapi3.Parameter{}

	all := append(openapi3.Parameters{}, pathParams...)
	all = append(all, opParams...)

	for _, paramRef := range all {
		param := resolveParameterRef(doc, paramRef)
		if param == nil {
			continue
		}
		key := param.In + ":" + param.Name
		merged[key] = param
	}

	return merged
}

func resolveParameterRef(doc *openapi3.T, ref *openapi3.ParameterRef) *openapi3.Parameter {
	if ref == nil {
		return nil
	}
	if ref.Value != nil {
		return ref.Value
	}
	if ref.Ref == "" || doc == nil || doc.Components == nil {
		return nil
	}

	const prefix = "#/components/parameters/"
	if !strings.HasPrefix(ref.Ref, prefix) {
		return nil
	}
	name := strings.TrimPrefix(ref.Ref, prefix)
	if doc.Components.Parameters == nil {
		return nil
	}
	paramRef, ok := doc.Components.Parameters[name]
	if !ok || paramRef == nil {
		return nil
	}
	return paramRef.Value
}

func extractPathParams(params map[string]*openapi3.Parameter, path string) []RawParam {
	matches := pathParamRegexp.FindAllStringSubmatch(path, -1)
	if len(matches) == 0 {
		return nil
	}

	var result []RawParam
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		name := match[1]
		param := params["path:"+name]
		if param == nil {
			result = append(result, RawParam{Name: name, In: "path", Required: true})
			continue
		}
		result = append(result, RawParam{
			Name:        param.Name,
			In:          param.In,
			Description: param.Description,
			Required:    param.Required,
			Schema:      schemaOrAny(param.Schema),
		})
	}

	return result
}

func extractQueryParams(params map[string]*openapi3.Parameter) []RawParam {
	var names []string
	for key := range params {
		if strings.HasPrefix(key, "query:") {
			names = append(names, strings.TrimPrefix(key, "query:"))
		}
	}
	sort.Strings(names)

	result := make([]RawParam, 0, len(names))
	for _, name := range names {
		param := params["query:"+name]
		if param == nil {
			continue
		}
		schema := schemaOrAny(param.Schema)
		if schema.Value != nil && schema.Value.Type != nil && schema.Value.Type.Is("array") && schema.Value.Items == nil {
			schema.Value.Items = &openapi3.SchemaRef{Value: &openapi3.Schema{}}
		}
		result = append(result, RawParam{
			Name:        param.Name,
			In:          param.In,
			Description: param.Description,
			Required:    param.Required,
			Schema:      schema,
		})
	}

	return result
}

func extractRequestBody(op *openapi3.Operation) (*RawBody, error) {
	if op == nil || op.RequestBody == nil {
		return nil, nil
	}

	body := op.RequestBody.Value
	if body == nil {
		return nil, nil
	}

	contentType, schemaRef := pickContentSchema(body.Content)
	if schemaRef == nil {
		return nil, nil
	}

	isForm := contentType == "multipart/form-data" || contentType == "application/x-www-form-urlencoded"

	return &RawBody{
		Schema:     schemaRef,
		Required:   body.Required,
		IsFormData: isForm,
	}, nil
}

func extractResponseSchema(op *openapi3.Operation) *openapi3.SchemaRef {
	if op == nil || op.Responses == nil {
		return nil
	}

	resp := pickResponse(op.Responses)
	if resp == nil {
		return nil
	}

	_, schemaRef := pickContentSchema(resp.Value.Content)
	return schemaRef
}

func pickResponse(responses *openapi3.Responses) *openapi3.ResponseRef {
	if responses == nil {
		return nil
	}

	if responses.Value("200") != nil {
		return responses.Value("200")
	}
	if responses.Value("201") != nil {
		return responses.Value("201")
	}

	var keys []string
	for key := range responses.Map() {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if strings.HasPrefix(key, "2") {
			return responses.Value(key)
		}
	}

	if responses.Default() != nil {
		return responses.Default()
	}

	return nil
}

func pickContentSchema(content openapi3.Content) (string, *openapi3.SchemaRef) {
	if content == nil {
		return "", nil
	}

	if mt, ok := content["application/json"]; ok && mt != nil {
		return "application/json", mt.Schema
	}
	if mt, ok := content["application/*+json"]; ok && mt != nil {
		return "application/*+json", mt.Schema
	}

	var keys []string
	for key := range content {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if len(keys) == 0 {
		return "", nil
	}
	mt := content[keys[0]]
	if mt == nil {
		return keys[0], nil
	}
	return keys[0], mt.Schema
}

func schemaOrAny(schema *openapi3.SchemaRef) *openapi3.SchemaRef {
	if schema != nil {
		return schema
	}
	return &openapi3.SchemaRef{Value: &openapi3.Schema{}}
}
