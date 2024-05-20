// This module provides the testable mock functions and variables for the environment variables. You can replace the functions and variables for testing.
//
// Providing mocks are:
//   - Environment variables
//   - Stdout, Stderr, and Stdin
//   - Exit and Args
//   - Password input function
package env

import (
	"os"
	"strconv"
)

// Env is the type that represents the environment variables.
type Env map[string]string

var (
	// DummyEnv is the map that stores dummy values for the environment variable.
	// This is used for testing.
	DummyEnv = make(Env)

	// Exit is the function that calls os.Exit normally.
	// If you want test, you can replace this function with a mock function.
	Exit = os.Exit

	// Args is the variable that stores the arguments of os.Args.
	// If you want test, you can replace this variable with a mock variable.
	Args = os.Args
)

// Getenv is the function that returns the value of the environment variable as a string.
// The value is searched in the DummyEnv first, and if it is not found, it is searched by os.Getenv.
// Finally, if the value is not found, it returns the defaultValue.
func Getenv(key string, defaultValue string) string {
	if DummyEnv[key] != "" {
		return DummyEnv[key]
	} else if value := os.Getenv(key); value != "" {
		return value
	} else {
		return defaultValue
	}
}

// SetenvDummy is the function that sets the DummyEnv.
// This is used for testing.
func SetenvDummy(key, value string) {
	DummyEnv[key] = value
}

// Getenv is the function that returns the value of the environment variable as a int.
// The value is searched in the DummyEnv first, and if it is not found, it is searched by os.Getenv.
// Finally, if the value is not found, it returns the defaultValue.
// If the value is not a number, it returns the defaultValue and an error.
func GetenvInt(key string, defaultValue int) (int, error) {
	var value string
	if DummyEnv[key] != "" {
		value = DummyEnv[key]
	} else if value = os.Getenv(key); value == "" {
		return defaultValue, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue, err
	}
	return intValue, nil
}

// Hostname is the function that returns the hostname of the machine.
// The value is searched "HOST" in the DummyEnv first, and if it is not found, it is returned by os.Hostname.
// Finally, if the value is not found (os.Hostname returns an error), it returns the defaultHost.
func Hostname(defaultHost string) string {
	if DummyEnv["HOST"] != "" {
		return DummyEnv["HOST"]
	}
	host, err := os.Hostname()
	if err != nil {
		return defaultHost
	}
	return host
}
