package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	var err error
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	zaplogger, err := config.Build()

	if err != nil {
		log.Fatalf("Cannot initialize zap logger: %v", err)
	}
	logger = zaplogger.Sugar()
	defer logger.Sync()
}

func Infof(message string, args ...interface{}) {
	logger.Infof(message, args)
}

func Debugf(message string, args ...interface{}) {
	logger.Debugf(message, args)
}

func Errorf(message string, args ...interface{}) {
	logger.Errorf(message, args)
}

func Fatalf(message string, args ...interface{}) {
	logger.Fatalf(message, args)
}

func Info(message string) {
	logger.Info(message)
}

func Debug(message string) {
	logger.Debug(message)
}

func Error(message string) {
	logger.Error(message)
}

func Fatal(message string) {
	logger.Fatal(message)
}