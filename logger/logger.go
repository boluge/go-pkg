package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	//LogLevel level
	LogLevel = DEBUG
	//LogInFile bool
	LogInFile = false
	//LogFile path
	LogFile = ""
)

const (
	//DEBUG LogLevel
	DEBUG = iota
	//INFO LogLevel
	INFO
	//WARN LogLevel
	WARN
	//ERROR LogLevel
	ERROR
)

func logfile() (*log.Logger, *os.File) {
	file, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(file, "", log.LstdFlags)
	return logger, file
}

// LogDebug debug
func LogDebug(v ...interface{}) {
	logger(DEBUG, v...)
}

// LogDebugf debug
func LogDebugf(source string, v ...interface{}) {
	loggerf(DEBUG, source, v...)
}

// Log info
func Log(v ...interface{}) {
	logger(INFO, v...)
}

// Logf info
func Logf(source string, v ...interface{}) {
	loggerf(INFO, source, v...)
}

// LogWarn warning
func LogWarn(v ...interface{}) {
	logger(WARN, v...)
}

// LogWarnf warning
func LogWarnf(source string, v ...interface{}) {
	loggerf(WARN, source, v...)
}

// LogError error
func LogError(v ...interface{}) {
	if v != nil {
		logger(ERROR, v...)
	}
}

// LogErrorf error
func LogErrorf(source string, v ...interface{}) {
	if v != nil {
		loggerf(ERROR, source, v...)
	}
}

// LogFatal Log error and exits
func LogFatal(v ...interface{}) {
	if v != nil {
		LogError(v...)
		os.Exit(1)
	}
}

// LogFatalf Log error and exits
func LogFatalf(source string, v ...interface{}) {
	if v != nil {
		LogErrorf(source, v...)
		os.Exit(1)
	}
}

// SetLevel to set log level
func SetLevel(level string) {
	switch level {
	case "DEBUG":
		LogLevel = DEBUG
	case "INFO":
		LogLevel = INFO
	case "WARN":
		LogLevel = WARN
	case "ERROR":
		LogLevel = ERROR
	default:
		LogLevel = ERROR
	}
}

// SetLogInFile to set if a log file is used
func SetLogInFile(b bool) {
	LogInFile = b
}

// SetLogFile to set the log file
func SetLogFile(file string) {
	LogFile = file
}

func logger(level int, v ...interface{}) {
	if LogLevel <= level && v != nil && v[0] != nil {
		v = append(v, "")
		copy(v[1:], v)

		switch level {
		case DEBUG:
			v[0] = "[DEBUG]"
		case INFO:
			v[0] = "[INFO]"
		case WARN:
			v[0] = "[WARN]"
		case ERROR:
			v[0] = "[ERROR]"
		}

		if LogInFile {
			logger, f := logfile()
			defer f.Close()
			logger.Println(v...)
		}

		log.Println(v...)
	}
}

func loggerf(level int, source string, v ...interface{}) {
	if LogLevel <= level && v != nil && v[0] != nil {

		switch level {
		case DEBUG:
			source = fmt.Sprintf("[DEBUG] %s", source)
		case INFO:
			source = fmt.Sprintf("[INFO] %s", source)
		case WARN:
			source = fmt.Sprintf("[WARN] %s", source)
		case ERROR:
			source = fmt.Sprintf("[ERROR] %s", source)
		}

		if LogInFile {
			logger, f := logfile()
			defer f.Close()
			logger.Println(v...)
		}

		log.Printf(source, v...)
	}
}
