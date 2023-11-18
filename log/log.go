package log

import (
	"log"
	"os"
)

type AgreeGateLoager struct {
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewAgreeGateLogger() *AgreeGateLoager {
	infoLogger := log.New(os.Stdout, "INFO", log.Flags())
	warnLogger := log.New(os.Stdout, "WARN", log.Flags())
	errorLogger := log.New(os.Stdout, "ERROR", log.Flags())
	logger := &AgreeGateLoager{}
	logger.InfoLogger = infoLogger
	logger.WarnLogger = warnLogger
	logger.ErrorLogger = errorLogger
	return logger
}

func (l *AgreeGateLoager) Info(v ...interface{}) {
	l.InfoLogger.Println(v...)
}

func (l *AgreeGateLoager) Warn(v ...interface{}) {
	l.WarnLogger.Println(v...)
}

func (l *AgreeGateLoager) Error(v ...interface{}) {
	l.ErrorLogger.Println(v...)
}
