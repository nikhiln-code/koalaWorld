package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init(isProduction bool) {
	if isProduction {
		// Create a lumberjack logger instance for log rotation
		lumberjackLogger := &lumberjack.Logger{
			Filename:   "/var/log/app/app.log", // Mount this path in Docker
			MaxSize:    10,  // MB
			MaxBackups: 5,
			MaxAge:     30,  // Days
			Compress:   true,
		}

		writeSyncer := zapcore.AddSync(lumberjackLogger)

		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder // human-readable timestamp

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			writeSyncer,
			zapcore.InfoLevel,
		)

		log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	} else {
		devLogger, err := zap.NewDevelopment()
		if err != nil {
			panic("failed to create development logger: " + err.Error())
		}
		log = devLogger
	}

	zap.ReplaceGlobals(log)
}

func Log() *zap.Logger {
	if log == nil {
		panic("logger not initialized: call logger.Init() before using logger")
	}
	return log
}

func SugarLog() *zap.SugaredLogger {
	return Log().Sugar()
}
