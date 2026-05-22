package utils

import "corekit/envutil"

// Deprecated: use corekit/envutil.Get.
// GetEnv gets an environment variable with a default value.
func GetEnv(key, defaultValue string) string {
	return envutil.Get(key, defaultValue)
}

// Deprecated: use corekit/envutil.GetInt.
// GetEnvInt gets an integer environment variable with a default value.
func GetEnvInt(key string, defaultValue int) int {
	return envutil.GetInt(key, defaultValue)
}

// Deprecated: use corekit/envutil.GetBool.
// GetEnvBool gets a boolean environment variable with a default value.
func GetEnvBool(key string, defaultValue bool) bool {
	return envutil.GetBool(key, defaultValue)
}
