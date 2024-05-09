package env

import (
	"testing"
)

func TestDummyPasswordInput(t *testing.T) {
	SetDummyPassword("dummy")
	if password, err := InputPassword(); err != nil {
		t.Error(err)
	} else if password != "dummy" {
		t.Error("Password is not correct")
	}
}

func TestDummyPasswordInputInterrupt(t *testing.T) {
	SetDummyPassword(Interrupt)
	if _, err := InputPassword(); err != ErrInterrupted {
		t.Error(err)
	}
}
