package env

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestESeries(t *testing.T) {
	err := os.RemoveAll("tmp/log")
	if err != nil {
		t.Fatal(err)
	}
	err = os.MkdirAll("tmp/log", 0755)
	if err != nil {
		t.Fatal(err)
	}
	SetLogger("tmp/log")

	buff := new(bytes.Buffer)
	Err = buff

	// EDebug("debug") with no debug.log
	SetDebug(true)
	EDebug("debug")
	str := buff.String()
	if str != "DEBUG: debug\n" {
		t.Error(str)
	}
	_, err = os.Stat("tmp/log/debug.log")
	if err == nil {
		t.Error(err)
	}
	buff.Reset()

	// EDebug("debug") with debug.log
	_, err = os.Create("tmp/log/debug.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("tmp/log/debug.log")
	EDebug("debug")
	str = buff.String()
	if str != "DEBUG: debug\n" {
		t.Error(str)
	}
	_, err = os.Stat("tmp/log/debug.log")
	if err != nil {
		t.Error(err)
	}
	log, err := os.ReadFile("tmp/log/debug.log")
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(string(log), "debug:debug") == false {
		t.Error(string(log))
	}
	buff.Reset()

	// EInfo("info")
	EInfo("info")
	str = buff.String()
	if str != "info\n" {
		t.Error(str)
	}
	_, err = os.Stat("tmp/log/info.log")
	if err != nil {
		t.Error(err)
	}
	log, err = os.ReadFile("tmp/log/info.log")
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(string(log), "info:info") == false {
		t.Error(string(log))
	}
	buff.Reset()

	// ENotice("notice")
	ENotice("notice")
	str = buff.String()
	if str != "notice\n" {
		t.Error(str)
	}
	_, err = os.Stat("tmp/log/notice.log")
	if err != nil {
		t.Error(err)
	}
	log, err = os.ReadFile("tmp/log/notice.log")
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(string(log), "notice:notice") == false {
		t.Error(string(log))
	}
	buff.Reset()

	// EError("error")
	EError("error")
	str = buff.String()
	if str != "error\n" {
		t.Error(str)
	}
	_, err = os.Stat("tmp/log/error.log")
	if err != nil {
		t.Error(err)
	}
	log, err = os.ReadFile("tmp/log/error.log")
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(string(log), "error:error") == false {
		t.Error(string(log))
	}
	buff.Reset()

}
