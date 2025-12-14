package main

import (
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *slog.Logger

func init() {

	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/api_log.log",
		MaxSize:    10,
		MaxAge:     1,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   false,
	}

	textHandler := slog.NewJSONHandler(lumberjackLogger, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelInfo,
		ReplaceAttr: nil,
	})
	Logger = slog.New(textHandler)

}
