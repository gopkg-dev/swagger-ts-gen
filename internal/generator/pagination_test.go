package generator

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestResolveReturnType_RecognizesPaginationAllOfOverride(t *testing.T) {
	components := openapi3.NewComponents()
	components.Schemas = openapi3.Schemas{
		"Response": {
			Value: &openapi3.Schema{
				Type: typesOf("object"),
				Properties: map[string]*openapi3.SchemaRef{
					"success": {Value: &openapi3.Schema{Type: typesOf("boolean")}},
					"data":    {Value: &openapi3.Schema{}},
				},
			},
		},
		"PaginationData": {
			Value: &openapi3.Schema{
				Type: typesOf("object"),
				Properties: map[string]*openapi3.SchemaRef{
					"count": {Value: &openapi3.Schema{Type: typesOf("integer")}},
					"list":  {Value: &openapi3.Schema{}},
				},
			},
		},
		"Api": {
			Value: &openapi3.Schema{
				Type:       typesOf("object"),
				Properties: map[string]*openapi3.SchemaRef{"id": {Value: &openapi3.Schema{Type: typesOf("string")}}},
			},
		},
	}
	doc := &openapi3.T{Components: &components}
	registry := NewTypeRegistry(doc)

	responseSchema := &openapi3.SchemaRef{Value: &openapi3.Schema{
		AllOf: openapi3.SchemaRefs{
			{Ref: "#/components/schemas/Response"},
			{Value: &openapi3.Schema{
				Type: typesOf("object"),
				Properties: map[string]*openapi3.SchemaRef{
					"data": {
						Value: &openapi3.Schema{
							AllOf: openapi3.SchemaRefs{
								{Ref: "#/components/schemas/PaginationData"},
								{Value: &openapi3.Schema{
									Type: typesOf("object"),
									Properties: map[string]*openapi3.SchemaRef{
										"list": {
											Value: &openapi3.Schema{
												Type:  typesOf("array"),
												Items: &openapi3.SchemaRef{Ref: "#/components/schemas/Api"},
											},
										},
									},
								}},
							},
						},
					},
				},
			}},
		},
	}}

	returnInfo, usedTypes := resolveReturnType("queryApis", responseSchema, registry, true)
	if returnInfo.Type != "PageResult<Api>" {
		t.Fatalf("unexpected return type: got %s want %s", returnInfo.Type, "PageResult<Api>")
	}
	if !returnInfo.UsesPageResult {
		t.Fatal("expected UsesPageResult=true")
	}
	if !reflect.DeepEqual(usedTypes, []string{"Api"}) {
		t.Fatalf("unexpected used types: got %v want %v", usedTypes, []string{"Api"})
	}
}

func TestResolveReturnType_DoesNotTreatAnyListAsPageResult(t *testing.T) {
	components := openapi3.NewComponents()
	components.Schemas = openapi3.Schemas{
		"PaginationData": {
			Value: &openapi3.Schema{
				Type: typesOf("object"),
				Properties: map[string]*openapi3.SchemaRef{
					"count": {Value: &openapi3.Schema{Type: typesOf("integer")}},
					"list":  {Value: &openapi3.Schema{}},
				},
			},
		},
	}
	doc := &openapi3.T{Components: &components}
	registry := NewTypeRegistry(doc)

	responseSchema := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: typesOf("object"),
		Properties: map[string]*openapi3.SchemaRef{
			"data": {Ref: "#/components/schemas/PaginationData"},
		},
	}}

	returnInfo, _ := resolveReturnType("queryAnything", responseSchema, registry, true)
	if returnInfo.UsesPageResult {
		t.Fatal("expected UsesPageResult=false when list is untyped any")
	}
	if returnInfo.Type != "PaginationData" {
		t.Fatalf("unexpected return type: got %s want %s", returnInfo.Type, "PaginationData")
	}
}
