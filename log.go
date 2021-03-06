package main

import (
	"fmt"
	"github.com/kataras/golog"
	"os"
	"path/filepath"
)

// Loggers

// AppLogger : logger
func AppLogger() (*golog.Logger, *os.File) {

	appLogger := golog.New()

	applogfile := appLogFile()

	if debugmode {
		appLogger.SetLevel("debug")
		appLogger.AddOutput(applogfile)
	} else {
		appLogger.SetLevel("warn")
		appLogger.SetOutput(applogfile)
	}

	return appLogger, applogfile

}

// GetLogger : logger
func GetLogger() (*golog.Logger, *os.File) {

	getLogger := golog.New()

	getlogfile := getLogFile()

	if debugmode {
		getLogger.SetLevel("debug")
		getLogger.AddOutput(getlogfile)
	} else {
		getLogger.SetLevel("warn")
		getLogger.SetOutput(getlogfile)
	}

	return getLogger, getlogfile

}

// PutLogger : logger
func PutLogger() (*golog.Logger, *os.File) {

	putLogger := golog.New()

	putlogfile := putLogFile()

	if debugmode {
		putLogger.SetLevel("debug")
		putLogger.AddOutput(putlogfile)
	} else {
		putLogger.SetLevel("warn")
		putLogger.SetOutput(putlogfile)
	}

	return putLogger, putlogfile

}

// DelLogger : logger
func DelLogger() (*golog.Logger, *os.File) {

	delLogger := golog.New()

	dellogfile := delLogFile()

	if debugmode {
		delLogger.SetLevel("debug")
		delLogger.AddOutput(dellogfile)
	} else {
		delLogger.SetLevel("warn")
		delLogger.SetOutput(dellogfile)
	}

	return delLogger, dellogfile

}

// Log Paths

func todayAppFilename() string {
	logfile := filepath.Clean(logdir + "/app.log")
	return logfile
}

func todayGetFilename() string {
	logfile := filepath.Clean(logdir + "/get.log")
	return logfile
}

func todayPutFilename() string {
	logfile := filepath.Clean(logdir + "/put.log")
	return logfile
}

func todayDelFilename() string {
	logfile := filepath.Clean(logdir + "/del.log")
	return logfile
}

// Log Files

func appLogFile() *os.File {

	filename := todayAppFilename()
	applogfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, logmode)
	if err != nil {
		fmt.Printf("Can`t open/create 'app' log file error | File [%s] | %v", filename, err)
		os.Exit(1)
	}

	err = os.Chmod(filename, logmode)
	if err != nil {
		fmt.Printf("Can`t chmod log file error | File [%s] | %v", filename, err)
		os.Exit(1)
	}

	return applogfile
}

func getLogFile() *os.File {
	filename := todayGetFilename()
	getlogfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, logmode)
	if err != nil {
		fmt.Printf("Can`t open/create 'get' log file error | File [%s] | %v", filename, err)
		os.Exit(1)
	}

	err = os.Chmod(filename, logmode)
	if err != nil {
		fmt.Printf("Can`t chmod log file error | File [%s] | %v", filename, err)
		os.Exit(1)
	}

	return getlogfile
}

func putLogFile() *os.File {
	filename := todayPutFilename()
	putlogfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, logmode)
	if err != nil {
		fmt.Printf("Can`t open/create 'put' log file error | File [%s] | %v", filename, err)
		os.Exit(1)
	}

	err = os.Chmod(filename, logmode)
	if err != nil {
		fmt.Printf("Can`t chmod log file error | File [%s] | %v", filename, err)
		os.Exit(1)
	}

	return putlogfile
}

func delLogFile() *os.File {
	filename := todayDelFilename()
	dellogfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, logmode)
	if err != nil {
		fmt.Printf("Can`t open/create 'del' log file error | File [%s] | %v", filename, err)
		os.Exit(1)
	}

	err = os.Chmod(filename, logmode)
	if err != nil {
		fmt.Printf("Can`t chmod log file error | File [%s] | %v", filename, err)
		os.Exit(1)
	}

	return dellogfile
}
