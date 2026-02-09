package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestGenerate_WritesRootIndexFile(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/api/v1/demo", &openapi3.PathItem{Get: &openapi3.Operation{Summary: "查询示例"}})

	outputDir := filepath.Join(t.TempDir(), "generated", "api")
	report, err := New(doc, Options{OutputDir: outputDir}).Generate()
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}
	if report == nil {
		t.Fatal("Generate returned nil report")
	}
	if report.Groups != 1 || report.Operations != 1 {
		t.Fatalf("unexpected report: %+v", *report)
	}

	indexPath := filepath.Join(outputDir, "index.ts")
	indexContent, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("read generated root index failed: %v", err)
	}

	if string(indexContent) != renderRootIndexFile() {
		t.Fatalf("unexpected root index content\n--- got ---\n%s\n--- want ---\n%s", string(indexContent), renderRootIndexFile())
	}
}

func TestGenerate_DeduplicatesCrossGroupModelsByReExport(t *testing.T) {
	doc := buildCrossGroupDuplicateModelDoc()

	outputDir := filepath.Join(t.TempDir(), "generated", "api")
	_, err := New(doc, Options{
		OutputDir:              outputDir,
		DedupeCrossGroupModels: true,
	}).Generate()
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	alphaModelPath := filepath.Join(outputDir, "alpha", "model", "index.ts")
	alphaModelContent, err := os.ReadFile(alphaModelPath)
	if err != nil {
		t.Fatalf("read alpha model failed: %v", err)
	}
	if !strings.Contains(string(alphaModelContent), "export interface User") {
		t.Fatalf("alpha model should define User:\n%s", string(alphaModelContent))
	}

	betaModelPath := filepath.Join(outputDir, "beta", "model", "index.ts")
	betaModelContent, err := os.ReadFile(betaModelPath)
	if err != nil {
		t.Fatalf("read beta model failed: %v", err)
	}
	if !strings.Contains(string(betaModelContent), "export type { User } from '../../alpha/model';") {
		t.Fatalf("beta model should re-export User from alpha model:\n%s", string(betaModelContent))
	}
	if strings.Contains(string(betaModelContent), "export interface User") {
		t.Fatalf("beta model should not redefine User:\n%s", string(betaModelContent))
	}
}

func TestGenerate_KeepDuplicateCrossGroupModelsByDefault(t *testing.T) {
	doc := buildCrossGroupDuplicateModelDoc()

	outputDir := filepath.Join(t.TempDir(), "generated", "api")
	_, err := New(doc, Options{OutputDir: outputDir}).Generate()
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	betaModelPath := filepath.Join(outputDir, "beta", "model", "index.ts")
	betaModelContent, err := os.ReadFile(betaModelPath)
	if err != nil {
		t.Fatalf("read beta model failed: %v", err)
	}
	if !strings.Contains(string(betaModelContent), "export interface User") {
		t.Fatalf("beta model should define User when dedupe option is disabled:\n%s", string(betaModelContent))
	}
	if strings.Contains(string(betaModelContent), "export type { User } from '../../alpha/model';") {
		t.Fatalf("beta model should not re-export User when dedupe option is disabled:\n%s", string(betaModelContent))
	}
}

func buildCrossGroupDuplicateModelDoc() *openapi3.T {
	components := openapi3.NewComponents()
	components.Schemas = openapi3.Schemas{
		"User": {
			Value: &openapi3.Schema{
				Type: typesOf("object"),
				Properties: openapi3.Schemas{
					"id": {Value: &openapi3.Schema{Type: typesOf("string")}},
				},
			},
		},
	}

	userDataSchema := &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Type: typesOf("object"),
			Properties: openapi3.Schemas{
				"data": {Ref: "#/components/schemas/User"},
			},
		},
	}
	successResponse := openapi3.NewResponse().
		WithDescription("ok").
		WithContent(openapi3.NewContentWithJSONSchemaRef(userDataSchema))

	doc := &openapi3.T{
		Components: &components,
		Paths:      openapi3.NewPaths(),
	}
	doc.Paths.Set("/api/v1/alpha", &openapi3.PathItem{
		Get: &openapi3.Operation{
			Summary:   "查询 alpha",
			Responses: openapi3.NewResponses(openapi3.WithStatus(200, &openapi3.ResponseRef{Value: successResponse})),
		},
	})
	doc.Paths.Set("/api/v1/beta", &openapi3.PathItem{
		Get: &openapi3.Operation{
			Summary:   "查询 beta",
			Responses: openapi3.NewResponses(openapi3.WithStatus(200, &openapi3.ResponseRef{Value: successResponse})),
		},
	})

	return doc
}
