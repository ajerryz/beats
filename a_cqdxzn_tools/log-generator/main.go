package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func main() {
	rollingLog := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    1,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   false,
	}
	logger := log.New(rollingLog, "log-generator", log.LstdFlags|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)

	logCount := 1000
	for i := 0; i < logCount; i++ {
		logger.Printf("helllo %d\n", i)
	}

}
