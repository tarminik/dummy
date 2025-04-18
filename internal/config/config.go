// Package config provides helpers for working with service configuration files.
// It helps list available configs and check if a config exists.
package config

import (
	"os"
	"strings"
)

// ListConfigs returns the list of service names (without .yaml extension) from the configs directory.
// It reads the given directory and collects all YAML config names.
func ListConfigs(configsDir string) ([]string, error) {
	files, err := os.ReadDir(configsDir)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			name := strings.TrimSuffix(file.Name(), ".yaml")
			names = append(names, name)
		}
	}
	return names, nil
}

// ConfigExists checks if a config file exists for a given service name.
// Returns true if the file configsDir/service.yaml exists, false otherwise.
func ConfigExists(configsDir, service string) bool {
	configPath := configsDir + "/" + service + ".yaml"
	_, err := os.Stat(configPath)
	return err == nil
}
