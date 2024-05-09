package env

import (
	"fmt"
	"io"
	"os"
)

var (
	// Out is the io.Writer. Default is os.Stdout.
	// If you want test, you can replace this variable with a mock variable.
	Out = io.Writer(os.Stdout)

	// Err is the io.Writer. Default is os.Stderr.
	// If you want test, you can replace this variable with a mock variable.
	Err = io.Writer(os.Stderr)

	// In is the io.Reader. Default is os.Stdin.
	// If you want test, you can replace this variable with a mock variable.
	In = io.Reader(os.Stdin)
)

// Outf is the function that outputs the formatted string to env.Out.
func Outf(format string, a ...any) {
	fmt.Fprintf(Out, format, a...)
}

// Errf is the function that outputs the formatted string to env.Err.
func Errf(format string, a ...any) {
	fmt.Fprintf(Err, format, a...)
}
