package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

// func NewLogger(writer io.Writer) *Logger {
// 	return &Logger{
// 		debug: log.New(writer, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
// 		info:  log.New(writer, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
// 		warn:  log.New(writer, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
// 		err:   log.New(writer, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
// 		writer: writer,
// 	}
// }

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)

	logger := log.New(
		writer,
		prefix,
		log.Ldate|log.Ltime,
	)

	return &Logger{
		debug:  log.New(writer, "DEBUG: ", logger.Flags()),
		info:   log.New(writer, "INFO: ", logger.Flags()),
		warn:   log.New(writer, "WARN: ", logger.Flags()),
		err:    log.New(writer, "ERROR: ", logger.Flags()),
		writer: writer,
	}
}

// Methods for the Logger struct
func (l *Logger) Debug(message string) {
	l.debug.Println(message)
}

func (l *Logger) Info(message string) {
	l.info.Println(message)
}

func (l *Logger) Warn(message string) {
	l.warn.Println(message)
}

func (l *Logger) Err(message string) {
	l.err.Println(message)
}

// Methods for the Logger struct with formatted strings
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
