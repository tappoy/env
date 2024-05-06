// This module provides the ability to retrieve environment variables with default values and types.
// It also provides the ability to set dummy values for ease of testing.
package env

import (
	"fmt"
	"io"
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

	// Out is the io.Writer. Default is os.Stdout.
	// If you want test, you can replace this variable with a mock variable.
	Out = io.Writer(os.Stdout)

	// Err is the io.Writer. Default is os.Stderr.
	// If you want test, you can replace this variable with a mock variable.
	Err = io.Writer(os.Stderr)

	// In is the io.Reader. Default is os.Stdin.
	// If you want test, you can replace this variable with a mock variable.
	In = io.Reader(os.Stdin)

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

// Outf is the function that outputs the formatted string to env.Out.
func Outf(format string, a ...any) {
	fmt.Fprintf(Out, format, a...)
}

// Errf is the function that outputs the formatted string to env.Err.
func Errf(format string, a ...any) {
	fmt.Fprintf(Err, format, a...)
}
