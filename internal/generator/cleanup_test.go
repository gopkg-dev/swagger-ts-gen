package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPruneStaleGroupDirs_RemovesStaleGeneratedDir(t *testing.T) {
	outputDir := t.TempDir()
	activeGroupDir := filepath.Join(outputDir, "sysApi")
	staleGroupDir := filepath.Join(outputDir, "sys-api")

	if err := os.MkdirAll(filepath.Join(activeGroupDir, "model"), 0o755); err != nil {
		t.Fatalf("create active group dir failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(activeGroupDir, "index.ts"), []byte("export {}"), 0o644); err != nil {
		t.Fatalf("write active index failed: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(staleGroupDir, "model"), 0o755); err != nil {
		t.Fatalf("create stale group dir failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(staleGroupDir, "index.ts"), []byte("export {}"), 0o644); err != nil {
		t.Fatalf("write stale index failed: %v", err)
	}

	if err := pruneStaleGroupDirs(outputDir, []string{"sysApi"}); err != nil {
		t.Fatalf("pruneStaleGroupDirs returned error: %v", err)
	}

	if _, err := os.Stat(activeGroupDir); err != nil {
		t.Fatalf("active group dir should remain: %v", err)
	}
	if _, err := os.Stat(staleGroupDir); !os.IsNotExist(err) {
		t.Fatalf("stale group dir should be removed, stat err=%v", err)
	}
}

func TestPruneStaleGroupDirs_IgnoreNonGeneratedDir(t *testing.T) {
	outputDir := t.TempDir()
	manualDir := filepath.Join(outputDir, "manual-not-generated")
	if err := os.MkdirAll(manualDir, 0o755); err != nil {
		t.Fatalf("create manual dir failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(manualDir, "README.md"), []byte("manual"), 0o644); err != nil {
		t.Fatalf("write manual file failed: %v", err)
	}

	if err := pruneStaleGroupDirs(outputDir, []string{"users"}); err != nil {
		t.Fatalf("pruneStaleGroupDirs returned error: %v", err)
	}

	if _, err := os.Stat(manualDir); err != nil {
		t.Fatalf("non-generated dir should remain: %v", err)
	}
}
