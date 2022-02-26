package tools

import "os"

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func EnvToString(value *string, key, defaultString string) {
	*value = getEnv(key, defaultString)
}
