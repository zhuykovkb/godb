package logger

import "go.uber.org/zap"

var (
	logger *zap.Logger
)

func Init() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger = log
}

func Info(s string) {
	logger.Info(s)
}

func Warn(s string) {
	logger.Warn(s)
}
