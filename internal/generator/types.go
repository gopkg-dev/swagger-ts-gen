package generator

import (
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type TypeDef struct {
	Name        string
	Schema      *openapi3.SchemaRef
	Description string
	Kind        string
	Extends     []string
}

type TypeRegistry struct {
	doc       *openapi3.T
	types     map[string]*TypeDef
	nameTaken map[string]bool
	refToName map[string]string
	typeOrder []string
}

func NewTypeRegistry(doc *openapi3.T) *TypeRegistry {
	return &TypeRegistry{
		doc:       doc,
		types:     map[string]*TypeDef{},
		nameTaken: map[string]bool{},
		refToName: map[string]string{},
	}
}

func (r *TypeRegistry) RegisterRef(ref string) (string, error) {
	if ref == "" {
		return "", fmt.Errorf("empty ref")
	}
	if name, ok := r.refToName[ref]; ok {
		return name, nil
	}

	name := nameFromRef(ref)
	if name == "" {
		return "", fmt.Errorf("invalid ref: %s", ref)
	}

	schemaRef := r.resolveRefSchema(ref)
	if schemaRef == nil {
		return "", fmt.Errorf("schema ref not found: %s", ref)
	}

	finalName := r.ensureUniqueName(name)
	r.refToName[ref] = finalName
	r.addType(&TypeDef{
		Name:   finalName,
		Schema: schemaRef,
		Kind:   "component",
	})

	return finalName, nil
}

func (r *TypeRegistry) RegisterInline(nameHint string, schema *openapi3.SchemaRef, description string) string {
	base := sanitizeTypeName(nameHint)
	name := r.ensureUniqueName(base)
	r.addType(&TypeDef{
		Name:        name,
		Schema:      schema,
		Description: description,
		Kind:        "inline",
	})
	return name
}

func (r *TypeRegistry) RegisterInlineWithExtends(nameHint string, schema *openapi3.SchemaRef, description string, extends []string) string {
	base := sanitizeTypeName(nameHint)
	name := r.ensureUniqueName(base)
	r.addType(&TypeDef{
		Name:        name,
		Schema:      schema,
		Description: description,
		Kind:        "inline",
		Extends:     uniqueStrings(extends),
	})
	return name
}

func (r *TypeRegistry) Types() []*TypeDef {
	names := append([]string{}, r.typeOrder...)
	sort.Strings(names)
	result := make([]*TypeDef, 0, len(names))
	for _, name := range names {
		def := r.types[name]
		if def != nil {
			result = append(result, def)
		}
	}
	return result
}

func (r *TypeRegistry) ensureUniqueName(base string) string {
	name := base
	idx := 2
	for r.nameTaken[name] {
		name = fmt.Sprintf("%s%d", base, idx)
		idx++
	}
	r.nameTaken[name] = true
	return name
}

func (r *TypeRegistry) addType(def *TypeDef) {
	if def == nil || def.Name == "" {
		return
	}
	if _, ok := r.types[def.Name]; ok {
		return
	}
	r.types[def.Name] = def
	r.typeOrder = append(r.typeOrder, def.Name)
}

func (r *TypeRegistry) resolveRefSchema(ref string) *openapi3.SchemaRef {
	if r.doc == nil || r.doc.Components == nil {
		return nil
	}
	if strings.HasPrefix(ref, "#/components/schemas/") {
		name := strings.TrimPrefix(ref, "#/components/schemas/")
		if r.doc.Components.Schemas == nil {
			return nil
		}
		if schemaRef, ok := r.doc.Components.Schemas[name]; ok {
			return schemaRef
		}
	}

	if strings.HasPrefix(ref, "#/definitions/") {
		name := strings.TrimPrefix(ref, "#/definitions/")
		if r.doc.Components.Schemas == nil {
			return nil
		}
		if schemaRef, ok := r.doc.Components.Schemas[name]; ok {
			return schemaRef
		}
	}

	return nil
}

func nameFromRef(ref string) string {
	if strings.HasPrefix(ref, "#/components/schemas/") {
		name := strings.TrimPrefix(ref, "#/components/schemas/")
		return simplifyTypeName(name)
	}
	if strings.HasPrefix(ref, "#/definitions/") {
		name := strings.TrimPrefix(ref, "#/definitions/")
		return simplifyTypeName(name)
	}
	return ""
}

func simplifyTypeName(raw string) string {
	if raw == "" {
		return ""
	}
	parts := strings.Split(raw, ".")
	last := parts[len(parts)-1]
	return sanitizeTypeName(last)
}

func (r *TypeRegistry) SchemaToType(schemaRef *openapi3.SchemaRef, deps map[string]struct{}) string {
	if schemaRef == nil {
		return "any"
	}
	if schemaRef.Ref != "" {
		name, err := r.RegisterRef(schemaRef.Ref)
		if err != nil {
			return "any"
		}
		if deps != nil {
			deps[name] = struct{}{}
		}
		return name
	}

	schema := schemaRef.Value
	if schema == nil {
		return "any"
	}

	base := r.schemaValueToType(schema, deps)
	if schema.Nullable {
		return base + " | null"
	}
	return base
}

func (r *TypeRegistry) schemaValueToType(schema *openapi3.Schema, deps map[string]struct{}) string {
	if schema == nil {
		return "any"
	}

	if len(schema.Enum) > 0 {
		return enumToType(schema.Enum)
	}

	if len(schema.OneOf) > 0 {
		return joinSchemaTypes(r, schema.OneOf, deps, " | ")
	}
	if len(schema.AnyOf) > 0 {
		return joinSchemaTypes(r, schema.AnyOf, deps, " | ")
	}
	if len(schema.AllOf) > 0 {
		return joinSchemaTypes(r, schema.AllOf, deps, " & ")
	}

	switch {
	case schema.Type != nil && schema.Type.Is("string"):
		if schema.Format == "binary" {
			return "Blob"
		}
		return "string"
	case schema.Type != nil && (schema.Type.Is("integer") || schema.Type.Is("number")):
		return "number"
	case schema.Type != nil && schema.Type.Is("boolean"):
		return "boolean"
	case schema.Type != nil && schema.Type.Is("array"):
		if schema.Items == nil {
			return "Array<any>"
		}
		itemType := r.SchemaToType(schema.Items, deps)
		return "Array<" + itemType + ">"
	case schema.Type != nil && schema.Type.Is("object"):
		return renderInlineObject(r, schema, deps, "")
	default:
		if len(schema.Properties) > 0 || schema.AdditionalProperties.Schema != nil || schema.AdditionalProperties.Has != nil {
			return renderInlineObject(r, schema, deps, "")
		}
	}

	return "any"
}

func renderInlineObject(r *TypeRegistry, schema *openapi3.Schema, deps map[string]struct{}, indent string) string {
	if schema == nil {
		return "Record<string, any>"
	}

	if len(schema.Properties) == 0 && schema.AdditionalProperties.Schema != nil {
		valueType := r.SchemaToType(schema.AdditionalProperties.Schema, deps)
		return "Record<string, " + valueType + ">"
	}

	if len(schema.Properties) == 0 {
		if schema.AdditionalProperties.Has != nil && !*schema.AdditionalProperties.Has {
			return "Record<string, never>"
		}
		return "Record<string, any>"
	}

	required := map[string]struct{}{}
	for _, name := range schema.Required {
		required[name] = struct{}{}
	}

	var keys []string
	for name := range schema.Properties {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString("{\n")
	for _, name := range keys {
		propSchema := schema.Properties[name]
		if propSchema == nil {
			continue
		}
		optional := "?"
		if _, ok := required[name]; ok {
			optional = ""
		}
		propName := name
		if !isValidIdentifier(name) {
			propName = fmt.Sprintf("'%s'", escapeTSString(name))
		}
		propType := r.SchemaToType(propSchema, deps)
		b.WriteString(indent + "  " + propName + optional + ": " + propType + ";\n")
	}
	b.WriteString(indent + "}")

	return b.String()
}

func enumToType(values []any) string {
	var parts []string
	for _, v := range values {
		switch val := v.(type) {
		case string:
			parts = append(parts, fmt.Sprintf("'%s'", escapeTSString(val)))
		case float64:
			parts = append(parts, fmt.Sprintf("%v", val))
		case int:
			parts = append(parts, fmt.Sprintf("%d", val))
		case bool:
			parts = append(parts, fmt.Sprintf("%t", val))
		default:
			parts = append(parts, "any")
		}
	}
	if len(parts) == 0 {
		return "any"
	}
	return strings.Join(parts, " | ")
}

func escapeTSString(value string) string {
	if value == "" {
		return value
	}
	escaped := strings.ReplaceAll(value, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "'", "\\'")
	escaped = strings.ReplaceAll(escaped, "\n", " ")
	return escaped
}

func joinSchemaTypes(r *TypeRegistry, refs openapi3.SchemaRefs, deps map[string]struct{}, sep string) string {
	var parts []string
	for _, ref := range refs {
		if ref == nil {
			continue
		}
		parts = append(parts, r.SchemaToType(ref, deps))
	}
	if len(parts) == 0 {
		return "any"
	}
	return strings.Join(parts, sep)
}
