package generator

import "testing"

func TestBuildQueryParamTypeName_UsesOperationNameSuffixParam(t *testing.T) {
	got := buildQueryParamTypeName("queryLoggers", "loggers")
	want := "queryLoggersParam"
	if got != want {
		t.Fatalf("unexpected name: got %s want %s", got, want)
	}
}

func TestBuildQueryParamTypeName_AvoidsDuplicateParamSuffix(t *testing.T) {
	got := buildQueryParamTypeName("queryLoggersParam", "loggers")
	want := "queryLoggersParam"
	if got != want {
		t.Fatalf("unexpected name: got %s want %s", got, want)
	}
}

func TestBuildQueryParamTypeName_FallsBackToGroupWhenOperationNameEmpty(t *testing.T) {
	got := buildQueryParamTypeName("", "loggers")
	want := "loggersParam"
	if got != want {
		t.Fatalf("unexpected name: got %s want %s", got, want)
	}
}
