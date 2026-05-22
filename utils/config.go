package utils

import "corekit/configutil"

// Deprecated: use corekit/configutil.ResolveConfFilePath.
// ResolveConfFilePath resolves the full path to a config file using SERVICE_HOME/conf/ as prefix.
func ResolveConfFilePath(configPath string) string {
	return configutil.ResolveConfFilePath(configPath)
}

// Deprecated: use corekit/configutil.LoadConfigMap.
// LoadConfigMap loads a YAML config file into a generic map.
// It preserves the previous behavior of returning nil on read/parse errors.
func LoadConfigMap(configPath string) map[string]any {
	return configutil.LoadConfigMap(configPath)
}

// Deprecated: use corekit/configutil.LoadConfigMapWithError.
// LoadConfigMapWithError loads a YAML config file into a generic map and returns errors directly.
func LoadConfigMapWithError(configPath string) (map[string]any, error) {
	return configutil.LoadConfigMapWithError(configPath)
}
