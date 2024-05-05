// Test env.go

package env

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	value := Getenv("TEST_KEY", "default_value")
	if value != "test_value" {
		t.Errorf("Getenv() = %s; want test_value", value)
	}
}

func TestGetenvWithDefaultValue(t *testing.T) {
	os.Setenv("TEST_KEY", "")
	value := Getenv("TEST_KEY", "default_value")
	if value != "default_value" {
		t.Errorf("Getenv() = %s; want default_value", value)
	}
}

func TestGetenvInt(t *testing.T) {
	os.Setenv("TEST_KEY", "100")
	value, _ := GetenvInt("TEST_KEY", 0)
	if value != 100 {
		t.Errorf("GetenvInt() = %d; want 100", value)
	}
}

func TestGetenvIntWithDefaultValue(t *testing.T) {
	os.Setenv("TEST_KEY", "")
	value, _ := GetenvInt("TEST_KEY", 0)
	if value != 0 {
		t.Errorf("GetenvInt() = %d; want 0", value)
	}
}

func TestGetenvIntWithInvalidValue(t *testing.T) {
	os.Setenv("TEST_KEY", "invalid")
	value, err := GetenvInt("TEST_KEY", 123)
	if err == nil {
		t.Errorf("GetenvInt() = %d; want error", value)
	}
	if value != 123 {
		t.Errorf("GetenvInt() = %d; want 123", value)
	}
}

func TestGetenvIntWithEmptyValue(t *testing.T) {
	os.Setenv("TEST_KEY", "")
	value, _ := GetenvInt("TEST_KEY", 123)
	if value != 123 {
		t.Errorf("GetenvInt() = %d; want 123", value)
	}
}

func TestGetenvWithDummyEnv(t *testing.T) {
	DummyEnv["TEST_KEY"] = "1st"
	os.Setenv("TEST_KEY", "2nd")
	value := Getenv("TEST_KEY", "3rd")
	if value != "1st" {
		t.Errorf("Getenv() = %s; want dummy_value", value)
	}
}

func TestGetenvIntWithDummyEnv(t *testing.T) {
	DummyEnv["TEST_KEY"] = "1"
	os.Setenv("TEST_KEY", "2")
	value, _ := GetenvInt("TEST_KEY", 3)
	if value != 1 {
		t.Errorf("GetenvInt() = %d; want 1", value)
	}
}
