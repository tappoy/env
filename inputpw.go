package env

import (
	"errors"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	// The input process is interrupted.
	ErrInterrupted = errors.New("ErrInterrupted")

	// The dummy password to generate ErrInterrupted.
	Interrupt = "Interrupt"
)

var dummyPassword *string

// InputPassword reads a password from the terminal.
// If the input is not from the terminal, it will read from the Env.In.
// If the dummy password is set, it will return the dummy password.
// If the dummy password is Interrupt, it will return ErrInterrupted.
//
// Errors:
//   - ErrInterrupted
//
// See:
//   - SetDummyPassword
//   - ClearDummyPassword
//   - Interrupt
func InputPassword() (string, error) {
	if dummyPassword != nil {
		return inputFromDummy(*dummyPassword)
	} else if terminal.IsTerminal(int(syscall.Stdin)) {
		return inputFromTerminal()
	} else {
		return inputFromReader(In)
	}
}

// SetDummyPassword is used to set the dummy password for testing.
func SetDummyPassword(pw string) {
	dummyPassword = &pw
}

// ClearDummyPassword is used to clear the dummy password for testing.
func ClearDummyPassword() {
	dummyPassword = nil
}

func inputFromDummy(pw string) (string, error) {
	if pw == Interrupt {
		return "", ErrInterrupted
	} else {
		return pw, nil
	}
}

func inputFromReader(reader io.Reader) (string, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func inputFromTerminal() (string, error) {
	// caputure the signal of Ctrl+C
	signalChan := make(chan os.Signal, 512)
	signal.Notify(signalChan, os.Interrupt)
	defer signal.Stop(signalChan)

	// create a channel to store the password
	pwChan := make(chan string)
	defer close(pwChan)

	// create a channel to store the error
	errChan := make(chan error)
	defer close(errChan)

	// copy the current terminal state
	currentState, err := terminal.GetState(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	go func() {
		<-signalChan
		// restore the terminal state after receiving Ctrl+C
		terminal.Restore(int(syscall.Stdin), currentState)
		errChan <- ErrInterrupted
	}()

	go func() {
		pw, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			pwChan <- ""
		} else {
			pwChan <- string(pw)
		}
	}()

	select {
	case pw := <-pwChan:
		return pw, nil
	case err := <-errChan:
		return "", err
	}
}
