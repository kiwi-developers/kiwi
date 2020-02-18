package log

import (
	"github.com/sirupsen/logrus"
	"github.com/xxxmailk/ylx/utils"
	"kiwi/manager/config"
	"log"
	"os"
)

var logHandle *logrus.Logger
var logEntry *logrus.Entry
var logFile *os.File
var logPath = "/var/log/kiwi/kiwi_manager.log"

func init() {
	SetLogger()
}

func handleLogFile() *os.File {
	if err := utils.CheckOrCreateDir("/var/log/kiwi"); err != nil {
		log.Fatalf("check log directory failed, %s", err)
	}
	fh, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("handle log file %s failed, %s", logPath, err)
	}
	return fh
}

// todo: close file handle before progress exited
func CloseLogFile() {
	if logFile != nil {
		err := logFile.Close()
		if err != nil {
			panic(err)
		}
	}
}

// create logger when progress loading
// you can invoke SetLogger when you want to reload config about log
func SetLogger() {
	logFile = handleLogFile()
	cfg := config.GetConf()
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{})
	l.SetReportCaller(true)
	l.SetLevel(parseLevel(cfg.Global.LogLevel))
	l.SetOutput(logFile)
	logHandle = l
	logEntry = logrus.NewEntry(logHandle)
}

// parse log level from string to logrus level
func parseLevel(level string) logrus.Level {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatalf("parse log level failed, %s", err)
	}
	return l
}

func GetLogger() *logrus.Logger {
	return logHandle
}

func GetLog() *logrus.Entry {
	return logEntry
}
