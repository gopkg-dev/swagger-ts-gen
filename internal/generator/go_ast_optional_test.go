package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseGoOptionalFieldsByType(t *testing.T) {
	rootDir := t.TempDir()
	schemaDir := filepath.Join(rootDir, "internal", "modules", "rbac", "schema")
	if mkdirErr := os.MkdirAll(schemaDir, 0o755); mkdirErr != nil {
		t.Fatalf("create schema dir failed: %v", mkdirErr)
	}
	goFilePath := filepath.Join(schemaDir, "model.go")
	goFileContent := `package sample

type Api struct {
	Path string ` + "`json:\"path\"`" + `
	Tag  string ` + "`json:\"tag,omitempty\"`" + `
	Skip string ` + "`json:\"-\"`" + `
}
`
	if writeErr := os.WriteFile(goFilePath, []byte(goFileContent), 0o644); writeErr != nil {
		t.Fatalf("write go file failed: %v", writeErr)
	}

	optionalFieldsByType, parseErr := ParseGoOptionalFieldsByType(rootDir, []string{"schema", "fiberx"})
	if parseErr != nil {
		t.Fatalf("ParseGoOptionalFieldsByType returned error: %v", parseErr)
	}

	candidates, exists := optionalFieldsByType["Api"]
	if !exists {
		t.Fatalf("Api optional fields not found: %+v", optionalFieldsByType)
	}
	if len(candidates) != 1 {
		t.Fatalf("unexpected candidate count: %+v", candidates)
	}
	apiOptionalFields := candidates[0].Fields
	fieldOrder := candidates[0].FieldOrder
	if apiOptionalFields["tag"] != true {
		t.Fatalf("tag should be optional: %+v", apiOptionalFields)
	}
	if apiOptionalFields["path"] != false {
		t.Fatalf("path should be required (not optional): %+v", apiOptionalFields)
	}
	if _, hasSkip := apiOptionalFields["-"]; hasSkip {
		t.Fatalf("ignored json field should not be included: %+v", apiOptionalFields)
	}
	if len(fieldOrder) != 2 || fieldOrder[0] != "path" || fieldOrder[1] != "tag" {
		t.Fatalf("unexpected field order: %+v", fieldOrder)
	}
}

func TestParseGoOptionalFieldsByType_CollectsDuplicateTypeNames(t *testing.T) {
	rootDir := t.TempDir()
	firstPackageDir := filepath.Join(rootDir, "internal", "modules", "rbac", "schema")
	secondPackageDir := filepath.Join(rootDir, "pkg", "fiberx")
	if mkdirErr := os.MkdirAll(firstPackageDir, 0o755); mkdirErr != nil {
		t.Fatalf("create first package dir failed: %v", mkdirErr)
	}
	if mkdirErr := os.MkdirAll(secondPackageDir, 0o755); mkdirErr != nil {
		t.Fatalf("create second package dir failed: %v", mkdirErr)
	}

	firstFile := filepath.Join(firstPackageDir, "model.go")
	secondFile := filepath.Join(secondPackageDir, "model.go")
	firstContent := `package a
type Api struct {
	Path string ` + "`json:\"path\"`" + `
}
`
	secondContent := `package b
type Api struct {
	Tag string ` + "`json:\"tag,omitempty\"`" + `
}
`
	if writeErr := os.WriteFile(firstFile, []byte(firstContent), 0o644); writeErr != nil {
		t.Fatalf("write first go file failed: %v", writeErr)
	}
	if writeErr := os.WriteFile(secondFile, []byte(secondContent), 0o644); writeErr != nil {
		t.Fatalf("write second go file failed: %v", writeErr)
	}

	optionalFieldsByType, parseErr := ParseGoOptionalFieldsByType(rootDir, []string{"schema", "fiberx"})
	if parseErr != nil {
		t.Fatalf("ParseGoOptionalFieldsByType returned error: %v", parseErr)
	}
	candidates, exists := optionalFieldsByType["Api"]
	if !exists {
		t.Fatalf("Api candidates should exist: %+v", optionalFieldsByType)
	}
	if len(candidates) != 2 {
		t.Fatalf("unexpected Api candidate count: got %d want 2", len(candidates))
	}
}

func TestParseGoOptionalFieldsByType_SkipsNonSchemaAndFiberx(t *testing.T) {
	rootDir := t.TempDir()
	ignoredDir := filepath.Join(rootDir, "internal", "modules", "rbac", "api")
	if mkdirErr := os.MkdirAll(ignoredDir, 0o755); mkdirErr != nil {
		t.Fatalf("create ignored dir failed: %v", mkdirErr)
	}
	goFilePath := filepath.Join(ignoredDir, "model.go")
	goFileContent := `package api
type Api struct {
	Path string ` + "`json:\"path,omitempty\"`" + `
}
`
	if writeErr := os.WriteFile(goFilePath, []byte(goFileContent), 0o644); writeErr != nil {
		t.Fatalf("write go file failed: %v", writeErr)
	}

	optionalFieldsByType, parseErr := ParseGoOptionalFieldsByType(rootDir, []string{"schema", "fiberx"})
	if parseErr != nil {
		t.Fatalf("ParseGoOptionalFieldsByType returned error: %v", parseErr)
	}
	if _, exists := optionalFieldsByType["Api"]; exists {
		t.Fatalf("non schema/fiberx file should not be parsed: %+v", optionalFieldsByType)
	}
}

func TestParseGoOptionalFieldsByType_UsesCustomIncludeDirs(t *testing.T) {
	rootDir := t.TempDir()
	customDir := filepath.Join(rootDir, "internal", "modules", "rbac", "dto")
	if mkdirErr := os.MkdirAll(customDir, 0o755); mkdirErr != nil {
		t.Fatalf("create custom dir failed: %v", mkdirErr)
	}
	goFilePath := filepath.Join(customDir, "model.go")
	goFileContent := `package dto
type Api struct {
	Path string ` + "`json:\"path,omitempty\"`" + `
}
`
	if writeErr := os.WriteFile(goFilePath, []byte(goFileContent), 0o644); writeErr != nil {
		t.Fatalf("write go file failed: %v", writeErr)
	}

	optionalFieldsByType, parseErr := ParseGoOptionalFieldsByType(rootDir, []string{"dto"})
	if parseErr != nil {
		t.Fatalf("ParseGoOptionalFieldsByType returned error: %v", parseErr)
	}
	candidates, exists := optionalFieldsByType["Api"]
	if !exists || len(candidates) != 1 {
		t.Fatalf("custom include dirs should be honored: %+v", optionalFieldsByType)
	}
}
