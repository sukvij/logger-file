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

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
)

func init() {
	file, err := os.OpenFile("myLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewAgreeGateLogger() *AgreeGateLoager {
	logger := &AgreeGateLoager{}
	logger.InfoLogger = InfoLog
	logger.WarnLogger = WarningLog
	logger.ErrorLogger = ErrorLog
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
