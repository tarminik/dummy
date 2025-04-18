package config

import (
	"os"
	"testing"
	"path/filepath"
)

func TestListConfigs(t *testing.T) {
	dir := t.TempDir()
	// Create some yaml and non-yaml files
	os.WriteFile(filepath.Join(dir, "service1.yaml"), []byte(""), 0644)
	os.WriteFile(filepath.Join(dir, "service2.yaml"), []byte(""), 0644)
	os.WriteFile(filepath.Join(dir, "not-a-config.txt"), []byte(""), 0644)

	services, err := ListConfigs(dir)
	if err != nil {
		t.Fatalf("ListConfigs error: %v", err)
	}
	if len(services) != 2 {
		t.Errorf("Expected 2 services, got %d", len(services))
	}
	found := map[string]bool{"service1": false, "service2": false}
	for _, s := range services {
		if _, ok := found[s]; ok {
			found[s] = true
		}
	}
	for name, ok := range found {
		if !ok {
			t.Errorf("Service %s not found in ListConfigs result", name)
		}
	}
}

func TestConfigExists(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "service1.yaml"), []byte(""), 0644)

	if !ConfigExists(dir, "service1") {
		t.Error("ConfigExists should return true for existing config")
	}
	if ConfigExists(dir, "nope") {
		t.Error("ConfigExists should return false for non-existing config")
	}
}
