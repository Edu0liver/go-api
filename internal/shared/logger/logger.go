package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	warn   *log.Logger
	info   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func NewLogger(w string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, w, log.Ldate|log.Ltime)

	return &Logger{
		debug:  log.New(writer, "DEBUG", logger.Flags()),
		warn:   log.New(writer, "WARN", logger.Flags()),
		info:   log.New(writer, "INFO", logger.Flags()),
		err:    log.New(writer, "ERR", logger.Flags()),
		writer: writer,
	}
}
