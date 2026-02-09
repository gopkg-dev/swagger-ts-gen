package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"
)

type GoStructOptionality struct {
	Fields     map[string]bool
	FieldOrder []string
}

// ParseGoOptionalFieldsByType parses Go struct definitions and returns optional json fields for each type.
// Rule: field is optional only when its json tag includes ",omitempty".
func ParseGoOptionalFieldsByType(rootDir string, includeDirs []string) (map[string][]GoStructOptionality, error) {
	trimmedRoot := strings.TrimSpace(rootDir)
	if trimmedRoot == "" {
		return nil, fmt.Errorf("root dir is empty")
	}
	includeDirSet := make(map[string]struct{}, len(includeDirs))
	for _, includeDir := range includeDirs {
		trimmedIncludeDir := strings.TrimSpace(includeDir)
		if trimmedIncludeDir == "" {
			continue
		}
		includeDirSet[trimmedIncludeDir] = struct{}{}
	}
	if len(includeDirSet) == 0 {
		includeDirSet["schema"] = struct{}{}
		includeDirSet["fiberx"] = struct{}{}
	}

	files := make([]string, 0, 256)
	err := filepath.WalkDir(trimmedRoot, func(path string, dirEntry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if dirEntry.IsDir() {
			name := dirEntry.Name()
			if name == ".git" || name == "vendor" || strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		if !shouldIncludeOptionalitySourceFile(trimmedRoot, path, includeDirSet) {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("scan go files failed: %w", err)
	}

	fileSet := token.NewFileSet()
	optionalFieldsByType := map[string][]GoStructOptionality{}

	for _, path := range files {
		parsedFile, parseErr := parser.ParseFile(fileSet, path, nil, parser.SkipObjectResolution)
		if parseErr != nil {
			return nil, fmt.Errorf("parse go file %s failed: %w", path, parseErr)
		}

		for _, declaration := range parsedFile.Decls {
			generalDeclaration, ok := declaration.(*ast.GenDecl)
			if !ok || generalDeclaration.Tok != token.TYPE {
				continue
			}
			for _, specification := range generalDeclaration.Specs {
				typeSpecification, ok := specification.(*ast.TypeSpec)
				if !ok {
					continue
				}
				structType, ok := typeSpecification.Type.(*ast.StructType)
				if !ok {
					continue
				}

				typeName := sanitizeTypeName(typeSpecification.Name.Name)
				optionalFields, fieldOrder := extractOptionalJSONFields(structType)
				optionalFieldsByType[typeName] = append(optionalFieldsByType[typeName], GoStructOptionality{
					Fields:     optionalFields,
					FieldOrder: fieldOrder,
				})
			}
		}
	}

	return optionalFieldsByType, nil
}

func shouldIncludeOptionalitySourceFile(rootDir string, path string, includeDirSet map[string]struct{}) bool {
	relativePath, err := filepath.Rel(rootDir, path)
	if err != nil {
		return false
	}
	segments := strings.Split(filepath.ToSlash(relativePath), "/")
	for _, segment := range segments {
		if _, exists := includeDirSet[segment]; exists {
			return true
		}
	}
	return false
}

func extractOptionalJSONFields(structType *ast.StructType) (map[string]bool, []string) {
	optionalFields := map[string]bool{}
	fieldOrder := make([]string, 0, 32)
	seenFieldOrder := map[string]struct{}{}
	if structType == nil || structType.Fields == nil {
		return optionalFields, fieldOrder
	}

	for _, field := range structType.Fields.List {
		if field == nil || len(field.Names) == 0 {
			continue
		}

		structTag := parseStructTag(field.Tag)
		for _, identifier := range field.Names {
			if identifier == nil || !identifier.IsExported() {
				continue
			}
			jsonName, isOptional, ignored := resolveJSONField(identifier.Name, structTag)
			if ignored || jsonName == "" {
				continue
			}
			optionalFields[jsonName] = isOptional
			if _, exists := seenFieldOrder[jsonName]; !exists {
				seenFieldOrder[jsonName] = struct{}{}
				fieldOrder = append(fieldOrder, jsonName)
			}
		}
	}

	return optionalFields, fieldOrder
}

func parseStructTag(tagLiteral *ast.BasicLit) reflect.StructTag {
	if tagLiteral == nil {
		return ""
	}
	rawTag := strings.Trim(tagLiteral.Value, "`")
	return reflect.StructTag(rawTag)
}

func resolveJSONField(fieldName string, structTag reflect.StructTag) (jsonName string, isOptional bool, ignored bool) {
	if fieldName == "" {
		return "", false, true
	}

	jsonTagValue := strings.TrimSpace(structTag.Get("json"))
	if jsonTagValue == "" {
		return fieldName, false, false
	}

	tagParts := strings.Split(jsonTagValue, ",")
	explicitName := strings.TrimSpace(tagParts[0])
	if explicitName == "-" {
		return "", false, true
	}

	fieldJSONName := explicitName
	if fieldJSONName == "" {
		fieldJSONName = fieldName
	}

	fieldOptional := false
	for _, option := range tagParts[1:] {
		if strings.TrimSpace(option) == "omitempty" {
			fieldOptional = true
			break
		}
	}

	return fieldJSONName, fieldOptional, false
}
