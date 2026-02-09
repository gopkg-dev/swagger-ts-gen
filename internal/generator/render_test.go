package generator

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRenderAPIFile_ModelImportsWithoutTrailingComma(t *testing.T) {
	content := RenderAPIFile(nil, []string{"SendEmailCodeForm", "Captcha", "GetCaptchaContentParam"}, false)

	expectedImportBlock := "import type {\n  Captcha,\n  GetCaptchaContentParam,\n  SendEmailCodeForm\n} from './model';\n"
	if !strings.Contains(content, expectedImportBlock) {
		t.Fatalf("unexpected model import block:\n%s", content)
	}

	unexpectedImportBlock := "SendEmailCodeForm,\n} from './model';"
	if strings.Contains(content, unexpectedImportBlock) {
		t.Fatalf("model import contains trailing comma before closing brace:\n%s", content)
	}
}

func TestRenderAPIFile_SingleModelImportWithoutTrailingComma(t *testing.T) {
	content := RenderAPIFile(nil, []string{"LoginForm"}, false)

	expectedImportBlock := "import type {\n  LoginForm\n} from './model';\n"
	if !strings.Contains(content, expectedImportBlock) {
		t.Fatalf("unexpected single model import block:\n%s", content)
	}

	unexpectedImportBlock := "LoginForm,\n} from './model';"
	if strings.Contains(content, unexpectedImportBlock) {
		t.Fatalf("single model import contains trailing comma before closing brace:\n%s", content)
	}
}

func TestRenderType_DefaultRequiredWithASTOptionality(t *testing.T) {
	registry := NewTypeRegistry(&openapi3.T{})
	registry.SetOptionalFieldsByType(map[string][]GoStructOptionality{
		"Api": {
			{
				Fields: map[string]bool{
					"tag": true,
				},
				FieldOrder: []string{"tag", "path"},
			},
		},
	})

	schemaRef := &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Type: typesOf("object"),
			Properties: openapi3.Schemas{
				"path": {Value: &openapi3.Schema{Type: typesOf("string")}},
				"tag":  {Value: &openapi3.Schema{Type: typesOf("string")}},
			},
		},
	}
	content, _ := RenderType(&TypeDef{Name: "Api", Schema: schemaRef}, registry)

	if !strings.Contains(content, "path: string;") {
		t.Fatalf("path should be required:\n%s", content)
	}
	if !strings.Contains(content, "tag?: string;") {
		t.Fatalf("tag should be optional:\n%s", content)
	}
}

func TestRenderType_DuplicateTypeNameSelectsBestASTCandidate(t *testing.T) {
	registry := NewTypeRegistry(&openapi3.T{})
	registry.SetOptionalFieldsByType(map[string][]GoStructOptionality{
		"Api": {
			{Fields: map[string]bool{}},
			{
				Fields: map[string]bool{
					"path":   false,
					"method": false,
					"tag":    true,
				},
				FieldOrder: []string{"method", "path", "tag"},
			},
		},
	})

	schemaRef := &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Type: typesOf("object"),
			Properties: openapi3.Schemas{
				"path":   {Value: &openapi3.Schema{Type: typesOf("string")}},
				"method": {Value: &openapi3.Schema{Type: typesOf("string")}},
				"tag":    {Value: &openapi3.Schema{Type: typesOf("string")}},
			},
		},
	}
	content, _ := RenderType(&TypeDef{Name: "Api", Schema: schemaRef}, registry)
	if !strings.Contains(content, "path: string;") || !strings.Contains(content, "method: string;") {
		t.Fatalf("path/method should be required:\n%s", content)
	}
	if !strings.Contains(content, "tag?: string;") {
		t.Fatalf("tag should be optional from best candidate:\n%s", content)
	}
}

func TestRenderType_UsesGoFieldOrderWhenMatched(t *testing.T) {
	registry := NewTypeRegistry(&openapi3.T{})
	registry.SetOptionalFieldsByType(map[string][]GoStructOptionality{
		"Api": {
			{
				Fields: map[string]bool{
					"path":   false,
					"method": false,
					"tag":    true,
				},
				FieldOrder: []string{"method", "path", "tag"},
			},
		},
	})

	schemaRef := &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Type: typesOf("object"),
			Properties: openapi3.Schemas{
				"path":   {Value: &openapi3.Schema{Type: typesOf("string")}},
				"method": {Value: &openapi3.Schema{Type: typesOf("string")}},
				"tag":    {Value: &openapi3.Schema{Type: typesOf("string")}},
			},
		},
	}
	content, _ := RenderType(&TypeDef{Name: "Api", Schema: schemaRef}, registry)
	methodIndex := strings.Index(content, "method: string;")
	pathIndex := strings.Index(content, "path: string;")
	tagIndex := strings.Index(content, "tag?: string;")
	if methodIndex == -1 || pathIndex == -1 || tagIndex == -1 {
		t.Fatalf("missing fields in content:\n%s", content)
	}
	if !(methodIndex < pathIndex && pathIndex < tagIndex) {
		t.Fatalf("unexpected field order, expected method -> path -> tag:\n%s", content)
	}
}
