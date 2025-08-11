package main

import (
	"flag"
	"log/slog"
)

type Options struct {
	FileName      string
	Level         string
	HttpLog       bool
	HttpGetLog    bool
	HttpPostLog   bool
	HttpPutLog    bool
	HttpDeleteLog bool
}

func NewOptions() *Options {
	return &Options{
		FileName:      "./logs/jsonloggenerator.log",
		Level:         "info",
		HttpLog:       true,
		HttpGetLog:    true,
		HttpPostLog:   true,
		HttpPutLog:    true,
		HttpDeleteLog: true,
	}
}

type Config struct {
	FileName      string
	Level         slog.Level
	HttpLog       bool
	HttpGetLog    bool
	HttpPostLog   bool
	HttpPutLog    bool
	HttpDeleteLog bool
}

func (opt Options) ParseFlag() *Config {
	flag.StringVar(&opt.FileName, "logFile", "./logs/jsonloggenerator.log", "log file name")
	flag.StringVar(&opt.Level, "logLevel", "info", "log level")
	flag.BoolVar(&opt.HttpLog, "logHttp", true, "http log")
	flag.BoolVar(&opt.HttpGetLog, "logGet", true, "http get log")
	flag.BoolVar(&opt.HttpPostLog, "logPost", true, "http post log")
	flag.BoolVar(&opt.HttpPutLog, "logPut", true, "http put log")
	flag.BoolVar(&opt.HttpDeleteLog, "logDelete", true, "http delete log")
	flag.Parse()

	var logLevel slog.Level
	switch opt.Level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	return &Config{
		FileName:      opt.FileName,
		Level:         logLevel,
		HttpLog:       opt.HttpLog,
		HttpGetLog:    opt.HttpGetLog,
		HttpPostLog:   opt.HttpPostLog,
		HttpPutLog:    opt.HttpPutLog,
		HttpDeleteLog: opt.HttpDeleteLog,
	}
}
