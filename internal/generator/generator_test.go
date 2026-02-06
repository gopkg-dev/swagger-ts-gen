package generator

import (
	"os"
	"path/filepath"
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
