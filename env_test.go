// Test main.go

package env

import (
  "os"
  "testing"
)

func TestGetEnvString(t *testing.T) {
  os.Setenv("TEST_KEY", "test_value")
  value := GetEnvString("TEST_KEY", "default_value")
  if value != "test_value" {
    t.Errorf("GetEnvString() = %s; want test_value", value)
  }
}

func TestGetEnvStringWithDefaultValue(t *testing.T) {
  os.Setenv("TEST_KEY", "")
  value := GetEnvString("TEST_KEY", "default_value")
  if value != "default_value" {
    t.Errorf("GetEnvString() = %s; want default_value", value)
  }
}

func TestGetEnvInt(t *testing.T) {
  os.Setenv("TEST_KEY", "100")
  value := GetEnvInt("TEST_KEY", 0)
  if value != 100 {
    t.Errorf("GetEnvInt() = %d; want 100", value)
  }
}

func TestGetEnvIntWithDefaultValue(t *testing.T) {
  os.Setenv("TEST_KEY", "")
  value := GetEnvInt("TEST_KEY", 0)
  if value != 0 {
    t.Errorf("GetEnvInt() = %d; want 0", value)
  }
}

func TestGetEnvIntWithInvalidValue(t *testing.T) {
  os.Setenv("TEST_KEY", "invalid")
  value := GetEnvInt("TEST_KEY", 0)
  if value != 0 {
    t.Errorf("GetEnvInt() = %d; want 0", value)
  }
}

func TestGetEnvIntWithEmptyValue(t *testing.T) {
  value := GetEnvInt("TEST_KEY", 0)
  if value != 0 {
    t.Errorf("GetEnvInt() = %d; want 0", value)
  }
}

func TestGetEnvIntWithNegativeValue(t *testing.T) {
  os.Setenv("TEST_KEY", "-100")
  value := GetEnvInt("TEST_KEY", 0)
  if value != -100 {
    t.Errorf("GetEnvInt() = %d; want -100", value)
  }
}

func TestGetEnvIntWithNegativeDefaultValue(t *testing.T) {
  value := GetEnvInt("TEST_KEY", -100)
  if value != -100 {
    t.Errorf("GetEnvInt() = %d; want -100", value)
  }
}

func TestGetEnvIntWithNegativeInvalidValue(t *testing.T) {
  os.Setenv("TEST_KEY", "invalid")
  value := GetEnvInt("TEST_KEY", -100)
  if value != -100 {
    t.Errorf("GetEnvInt() = %d; want -100", value)
  }
}
