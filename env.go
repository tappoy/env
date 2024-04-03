package env

import (
  "os"
  "strconv"
)

// GetEnvString returns the value of the environment variable key if it exists, otherwise it returns the defaultValue.
func GetEnvString(key string, defaultValue string) string {
  value := os.Getenv(key)
  if value == "" {
    return defaultValue
  }
  return value
}

// GetEnvInt returns the value of the environment variable key if it exists, otherwise it returns the defaultValue.
func GetEnvInt(key string, defaultValue int) int {
  value := os.Getenv(key)
  if value == "" {
    return defaultValue
  }
  intValue, err := strconv.Atoi(value)
  if err != nil {
    return defaultValue
  }
  return intValue
}
