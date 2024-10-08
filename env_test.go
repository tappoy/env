// Test env.go

package env

import (
	"bytes"
	"os"
	"testing"
)

func initDummyEnv() {
	SetenvDummy("TEST_KEY", "")
}

// Getenv
func TestGetenv(t *testing.T) {
	initDummyEnv()
	os.Setenv("TEST_KEY", "test_value")
	value := Getenv("TEST_KEY", "default_value")
	if value != "test_value" {
		t.Errorf("Getenv() = %s; want test_value", value)
	}
}

func TestGetenvWithDefaultValue(t *testing.T) {
	initDummyEnv()
	os.Setenv("TEST_KEY", "")
	value := Getenv("TEST_KEY", "default_value")
	if value != "default_value" {
		t.Errorf("Getenv() = %s; want default_value", value)
	}
}

func TestGetenvWithDummyEnv(t *testing.T) {
	initDummyEnv()
	SetenvDummy("TEST_KEY", "1st")
	os.Setenv("TEST_KEY", "2nd")
	value := Getenv("TEST_KEY", "3rd")
	if value != "1st" {
		t.Errorf("Getenv() = %s; want dummy_value", value)
	}
}

// GetenvInt
func TestGetenvInt(t *testing.T) {
	initDummyEnv()
	os.Setenv("TEST_KEY", "100")
	value, err := GetenvInt("TEST_KEY", 0)
	if err != nil {
		t.Errorf("GetenvInt() = %v; value: %v, want 100", err, value)
	}
	if value != 100 {
		t.Errorf("GetenvInt() = %d; want 100", value)
	}
}

func TestGetenvIntWithDefaultValue(t *testing.T) {
	initDummyEnv()
	os.Setenv("TEST_KEY", "")
	value, _ := GetenvInt("TEST_KEY", 0)
	if value != 0 {
		t.Errorf("GetenvInt() = %d; want 0", value)
	}
}

func TestGetenvIntWithInvalidValue(t *testing.T) {
	initDummyEnv()
	os.Setenv("TEST_KEY", "invalid")
	value, err := GetenvInt("TEST_KEY", 123)
	if err == nil {
		t.Errorf("GetenvInt() = %d; want error", value)
	}
	if value != 123 {
		t.Errorf("GetenvInt() = %d; want 123", value)
	}
}

func TestGetenvIntWithDummyEnv(t *testing.T) {
	initDummyEnv()
	SetenvDummy("TEST_KEY", "1")
	os.Setenv("TEST_KEY", "2")
	value, _ := GetenvInt("TEST_KEY", 3)
	if value != 1 {
		t.Errorf("GetenvInt() = %d; want 1", value)
	}
}

// GetenvBool
func TestGetenvBool(t *testing.T) {
	initDummyEnv()
	os.Setenv("TEST_KEY", "true")
	value, _ := GetenvBool("TEST_KEY", false)
	if value != true {
		t.Errorf("GetenvBool() = %v; want true", value)
	}
}

func TestGetenvBoolWithDefaultValue(t *testing.T) {
	initDummyEnv()
	os.Setenv("TEST_KEY", "")
	value, _ := GetenvBool("TEST_KEY", true)
	if value != true {
		t.Errorf("GetenvBool() = %v; want true", value)
	}
}

func TestGetenvBoolWithInvalidValue(t *testing.T) {
	initDummyEnv()
	os.Setenv("TEST_KEY", "invalid")
	value, err := GetenvBool("TEST_KEY", true)
	if err == nil {
		t.Errorf("GetenvBool() = %v; want error", value)
	}
	if value != true {
		t.Errorf("GetenvBool() = %v; want true", value)
	}
}

func TestGetenvBoolWithDummyEnv(t *testing.T) {
	initDummyEnv()
	SetenvDummy("TEST_KEY", "true")
	os.Setenv("TEST_KEY", "false")
	value, _ := GetenvBool("TEST_KEY", false)
	if value != true {
		t.Errorf("GetenvBool() = %v; want true", value)
	}
}

// Outf, Errf
func TestOutf(t *testing.T) {
	initDummyEnv()
	// set Out to bytes.Buffer
	Out = new(bytes.Buffer)
	Outf("test %s\n", "message")
	if Out.(*bytes.Buffer).String() != "test message\n" {
		t.Errorf("got<%s> want<test message>", Out.(*bytes.Buffer).String())
	}
}

func TestErrf(t *testing.T) {
	initDummyEnv()
	// set Err to bytes.Buffer
	Err = new(bytes.Buffer)
	Errf("test %s\n", "message")
	if Err.(*bytes.Buffer).String() != "test message\n" {
		t.Errorf("got<%s> want<test message>", Err.(*bytes.Buffer).String())
	}
}

// Hostname
func TestHostname(t *testing.T) {
	initDummyEnv()
	// set DummyEnv
	SetenvDummy("HOST", "testhost")
	host := Hostname("default_host")
	if host != "testhost" {
		t.Errorf("Hostname() = %s; want testhost", host)
	}

	// example of os.Hostname
	SetenvDummy("HOST", "")
	host = Hostname("default_host")
	t.Log(host)
}
