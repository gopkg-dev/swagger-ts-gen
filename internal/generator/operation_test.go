package generator

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractOperations_UsesDeterministicMethodOrder(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/api/v1/demo", &openapi3.PathItem{
		Post:    &openapi3.Operation{},
		Get:     &openapi3.Operation{},
		Delete:  &openapi3.Operation{},
		Put:     &openapi3.Operation{},
		Patch:   &openapi3.Operation{},
		Head:    &openapi3.Operation{},
		Options: &openapi3.Operation{},
	})

	expectedMethods := []string{"get", "post", "put", "patch", "delete", "head", "options"}
	for i := 0; i < 10; i++ {
		ops, err := ExtractOperations(doc)
		if err != nil {
			t.Fatalf("ExtractOperations returned error: %v", err)
		}

		gotMethods := make([]string, 0, len(ops))
		for _, operation := range ops {
			gotMethods = append(gotMethods, operation.Method)
		}

		if !reflect.DeepEqual(gotMethods, expectedMethods) {
			t.Fatalf("unexpected method order on run %d: got %v want %v", i+1, gotMethods, expectedMethods)
		}
	}
}

func TestExtractOperations_SortsPathsBeforeMethods(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/z", &openapi3.PathItem{Get: &openapi3.Operation{}})
	doc.Paths.Set("/a", &openapi3.PathItem{Post: &openapi3.Operation{}, Get: &openapi3.Operation{}})

	ops, err := ExtractOperations(doc)
	if err != nil {
		t.Fatalf("ExtractOperations returned error: %v", err)
	}

	gotOrder := make([]string, 0, len(ops))
	for _, operation := range ops {
		gotOrder = append(gotOrder, operation.Path+"#"+operation.Method)
	}

	wantOrder := []string{"/a#get", "/a#post", "/z#get"}
	if !reflect.DeepEqual(gotOrder, wantOrder) {
		t.Fatalf("unexpected operation order: got %v want %v", gotOrder, wantOrder)
	}
}
