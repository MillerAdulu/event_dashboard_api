package logging

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type LoggingLevel int

//go:generate stringer -type=LoggingLevel

const (
	TRACE LoggingLevel = iota
	INFO
	WARNING
	ERROR
	OFF
)

const _LoggingLevel_name = "TRACEINFOWARNINGERROROFF"

var _LoggingLevel_index = [...]uint8{0, 5, 9, 16, 21, 24}

type Logger struct {
	Name    string
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	level   LoggingLevel
}

var loggers = make(map[string]Logger)

func GetLogger(name string) Logger {
	return New(name, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
}

func (logger Logger) SetLevel(level LoggingLevel) Logger {
	switch level {
	case TRACE:
		logger.Trace.SetOutput(os.Stdout)
		logger.Info.SetOutput(os.Stdout)
		logger.Warning.SetOutput(os.Stdout)
		logger.Error.SetOutput(os.Stderr)
	case INFO:
		logger.Trace.SetOutput(ioutil.Discard)
		logger.Info.SetOutput(os.Stdout)
		logger.Warning.SetOutput(os.Stdout)
		logger.Error.SetOutput(os.Stderr)
	case WARNING:
		logger.Trace.SetOutput(ioutil.Discard)
		logger.Info.SetOutput(ioutil.Discard)
		logger.Warning.SetOutput(os.Stdout)
		logger.Error.SetOutput(os.Stderr)
	case ERROR:
		logger.Trace.SetOutput(ioutil.Discard)
		logger.Info.SetOutput(ioutil.Discard)
		logger.Warning.SetOutput(ioutil.Discard)
		logger.Error.SetOutput(os.Stderr)
	case OFF:
		logger.Trace.SetOutput(ioutil.Discard)
		logger.Info.SetOutput(ioutil.Discard)
		logger.Warning.SetOutput(ioutil.Discard)
		logger.Error.SetOutput(ioutil.Discard)
	}
	return logger
}

func (logger Logger) GetLevel() LoggingLevel {
	return logger.level
}

func New(
	name string,
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) Logger {
	loggers[name] = Logger{
		Name: name,
		Trace: log.New(traceHandle,
			"TRACE: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Info: log.New(infoHandle,
			"INFO: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(warningHandle,
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(errorHandle,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile),
	}
	return loggers[name]
}

func (i LoggingLevel) String() string {
	if i < 0 || i >= LoggingLevel(len(_LoggingLevel_index)-1) {
		return "LoggingLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LoggingLevel_name[_LoggingLevel_index[i]:_LoggingLevel_index[i+1]]
}
