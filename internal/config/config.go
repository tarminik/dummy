package config

import (
	"os"
	"strings"
)

// ListConfigs returns the list of service names (without .yaml) from the configs directory.
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
func ConfigExists(configsDir, service string) bool {
	configPath := configsDir + "/" + service + ".yaml"
	_, err := os.Stat(configPath)
	return err == nil
}
