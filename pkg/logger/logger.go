package logger

import "go.uber.org/zap"

var zapLogger *zap.Logger
var sugar *zap.SugaredLogger

func Info(msg string, keysAndValues ...any) {
	sugar.Infow(msg, keysAndValues...)
}

func InitLogger() error {
	var err error
	zapLogger, err = zap.NewProduction()
	sugar = zapLogger.Sugar()
	return err
}
