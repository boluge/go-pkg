package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	//LogLevel level
	LogLevel = INFO
	//LogInFile bool
	LogInFile = false
	//LogFile name
	LogFile = "application.log"
	//LogPathFile path file
	LogPathFile = "./application.log"
	//LogInTerm bool
	LogInTerm = true
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
	file, err := os.OpenFile(LogPathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(file, "", log.LstdFlags)
	return logger, file
}

// Debug debug
func Debug(v ...interface{}) {
	logger(DEBUG, v...)
}

// Debugf debug
func Debugf(source string, v ...interface{}) {
	loggerf(DEBUG, source, v...)
}

// Info
func Info(v ...interface{}) {
	logger(INFO, v...)
}

// Iinfof
func Infof(source string, v ...interface{}) {
	loggerf(INFO, source, v...)
}

// Warn warning
func Warn(v ...interface{}) {
	logger(WARN, v...)
}

// Warnf warning
func Warnf(source string, v ...interface{}) {
	loggerf(WARN, source, v...)
}

// Error error
func Error(v ...interface{}) {
	if v != nil {
		logger(ERROR, v...)
	}
}

// Errorf error
func Errorf(source string, v ...interface{}) {
	if v != nil {
		loggerf(ERROR, source, v...)
	}
}

// Fatal Log error and exits
func Fatal(err error) {
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
}

// Fatalf Log error and exits
func Fatalf(source string, v ...interface{}) {
	if v != nil {
		LogErrorf(source, v...)
		os.Exit(1)
	}
}

// SetLoggerConf func to set logger
func SetLoggerConf(level, file, output string, inFile, inTerm bool) {
	SetLevel(level)
	SetLogInTerm(inTerm)
	SetLogInFile(inFile)
	SetLogFile(output, file)
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

// SetLogInTerm to set if a log term is used
func SetLogInTerm(b bool) {
	LogInTerm = b
}

// SetLogInFile to set if a log file is used
func SetLogInFile(b bool) {
	LogInFile = b
}

// SetLogFile to set the log file
func SetLogFile(path, file string) {
	if len(file) > 0 {
		LogPathFile = fmt.Sprintf("%s/%s", path, file)
	}
	LogPathFile = fmt.Sprintf("%s/%s", path, LogFile)
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

		if LogInTerm {
			log.Println(v...)
		}
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
			logger.Printf(source, v...)
		}

		if LogInTerm {
			log.Printf(source, v...)
		}
	}
}
