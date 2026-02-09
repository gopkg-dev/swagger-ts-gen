package generator

import "testing"

func TestGroupFromPath_UsesLowerCamelForHyphenSegments(t *testing.T) {
	got := groupFromPath("/api/v1/sys-api")
	want := "sysApi"
	if got != want {
		t.Fatalf("unexpected group name: got %s want %s", got, want)
	}
}

func TestGroupFromPath_UsesLowerCamelForFirstSegment(t *testing.T) {
	got := groupFromPath("/dict-items/list")
	want := "dictItems"
	if got != want {
		t.Fatalf("unexpected group name: got %s want %s", got, want)
	}
}

func TestGroupFromPath_PrefixesNumericSegment(t *testing.T) {
	got := groupFromPath("/api/v1/123-system")
	want := "group123System"
	if got != want {
		t.Fatalf("unexpected group name: got %s want %s", got, want)
	}
}
