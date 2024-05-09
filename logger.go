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

// IsSetLogger is used to check if logger is set
func IsSetLogger() bool {
	return log != nil
}

// Debug is used to log debug message
func Debug(format string, args ...any) {
	if log == nil {
		return
	}
	log.Debug(format, args...)
}

// Error is used to log error message
func Error(format string, args ...any) {
	if log == nil {
		return
	}
	log.Error(format, args...)
}

// GetLogDir is used to get log directory
func GetLogDir() string {
	if log == nil {
		return ""
	}
	return log.GetLogDir()
}

// Info is used to log info message
func Info(format string, args ...any) {
	if log == nil {
		return
	}
	log.Info(format, args...)
}

// Notice is used to log notice message
func Notice(format string, args ...any) {
	if log == nil {
		return
	}
	log.Notice(format, args...)
}
