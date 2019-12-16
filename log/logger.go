package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
)

func init() {
	loggerConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: nil,
		InitialFields:    nil,
	}
	var err error
	Log, err = loggerConfig.Build()
	if err != nil {
		panic(err)
	}
}
