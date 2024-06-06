package main

import (
	"os"
	"strconv"
)

// GetAsString reads an environment variable or returns a default value
func GetAsString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetAsInt reads an environment variable into integer or returns a default value
func GetAsInt(key string, defaultValue int) int {
	valueStr := GetAsString(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
