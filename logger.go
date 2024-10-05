package env

import "github.com/tappoy/logger"

var log *logger.Logger

// SetLogger is used to set logger
// More info: See https://github.com/tappoy/logger
//
// Errors:
//   - logger.ErrCannotCreateLogDir
//   - logger.ErrCannotWriteLogFile
func SetLogger(logDir string) error {
	l, err := logger.NewLogger(logDir)
	if err != nil {
		return err
	}
	log = l
	return nil
}

// IsSetLogger is used to check if logger is set.
func IsSetLogger() bool {
	return log != nil
}

// GetLogDir is used to get log directory.
func GetLogDir() string {
	if log == nil {
		return ""
	}
	return log.GetLogDir()
}

var isDebug = false

// SetDebug is used to set debug mode.
// If debug mode is enabled, debug message will be output to env.Err with "DEBUG: " prefix in EDegug function.
func SetDebug(debug bool) {
	isDebug = debug
}

// GetDebug is used to get debug mode.
func GetDebug() bool {
	return isDebug
}

// Debug is used to log debug message.
func Debug(format string, args ...any) {
	if log == nil {
		return
	}
	log.Debug(format, args...)
}

// EDebug is used to log debug message and output to stderr.
// "\n" will be added to the end of the format.
func EDebug(format string, args ...any) {
	if isDebug {
		Errf("DEBUG: "+format+"\n", args...)
	}
	Debug(format, args...)
}

// Error is used to log error message.
func Error(format string, args ...any) {
	if log == nil {
		return
	}
	log.Error(format, args...)
}

// EError is used to log error message and output to stderr.
// "\n" will be added to the end of the format.
func EError(format string, args ...any) {
	Errf(format+"\n", args...)
	Error(format, args...)
}

// Info is used to log info message.
func Info(format string, args ...any) {
	if log == nil {
		return
	}
	log.Info(format, args...)
}

// EInfo is used to log info message and output to stderr.
// "\n" will be added to the end of the format.
func EInfo(format string, args ...any) {
	Errf(format+"\n", args...)
	Info(format, args...)
}

// Notice is used to log notice message.
func Notice(format string, args ...any) {
	if log == nil {
		return
	}
	log.Notice(format, args...)
}

// ENotice is used to log notice message and output to stderr.
// "\n" will be added to the end of the format.
func ENotice(format string, args ...any) {
	Errf(format+"\n", args...)
	Notice(format, args...)
}
