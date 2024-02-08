package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

import (
	"github.com/jeanphorn/log4go"
)

var Logger log4go.Logger
var logInitFlag bool

//if the directory of log not exist,then create it
func logDirInit(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

//generate the real log name
//*.log 	->	normal log
//*.wf.log 	->	error log
func logNameGen(spiderName string, dir string, typeIsOfEr bool) string {
	strings.TrimSuffix(dir, "/")

	var fileName string
	if typeIsOfEr {
		fileName = filepath.Join(dir, spiderName+"wf.log")
	} else {
		fileName = filepath.Join(dir, spiderName+".log")
	}

	return fileName
}

//trans the value of string to log4go.LevelType
func transToLogLevel(str string) log4go.Level {
	var logLevel log4go.Level

	str = strings.ToUpper(str)

	switch str {
	case "DEBUG":
		logLevel = log4go.DEBUG
	case "TRACE":
		logLevel = log4go.TRACE
	case "INFO":
		logLevel = log4go.INFO
	case "WARNING":
		logLevel = log4go.WARNING
	case "ERROR":
		logLevel = log4go.ERROR
	case "CRITICAL":
		logLevel = log4go.CRITICAL
	default:
		logLevel = log4go.INFO
	}
	return logLevel
}

//initial the global object Logger with parameters
// toStdOut 		-> whether print logs to stdout
// timeStr 			-> time of rolling over the logs, mostly choose MIDNIGHT(roll over at midnight)
// logRotateNum		-> logs rotation numbers

func InitialLogger(spiderName string, logLevel string, logDir string, toStdOut bool, timeStr string, logRotateNum int) error {
	if logInitFlag {
		return errors.New("Logger already initialized")
	}
	/*

		if !log4go(timeStr) {
			return fmt.Errorf("invalid value of when: %s", timeStr)
		}

	*/

	if err := logDirInit(logDir); err != nil {
		log4go.Error("create log dir failed logDirInit(%s)", logDir)
		return err
	}

	level := transToLogLevel(logLevel)

	//create the Logger
	Logger = make(log4go.Logger)
	if toStdOut {
		Logger.AddFilter("stdout", level, log4go.NewConsoleLogWriter())
	}
	//generate log name for the non-error logs
	fileName := logNameGen(spiderName, logDir, false)
	//create and initial a log writter for the non-error logs, no compress
	logWriter := log4go.NewFileLogWriter(fileName, true, true)
	if logWriter == nil {
		return fmt.Errorf("error in log4go.NewTimeFileLogWriter(%s)", fileName)
	}
	logWriter.SetFormat(log4go.FORMAT_DEFAULT)
	Logger.AddFilter("log", level, logWriter)

	fileNameWf := logNameGen(spiderName, logDir, true)
	logWriter = log4go.NewFileLogWriter(fileName, true, true)
	if logWriter == nil {
		return fmt.Errorf("error in log4go.NewTimeFileLogWriter(%s)", fileNameWf)
	}
	logWriter.SetFormat(log4go.FORMAT_DEFAULT)
	Logger.AddFilter("log_wf", log4go.WARNING, logWriter)

	logInitFlag = true
	return nil

}
