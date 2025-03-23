package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func SetupLogger() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	Logger = logger.Sugar()
}
