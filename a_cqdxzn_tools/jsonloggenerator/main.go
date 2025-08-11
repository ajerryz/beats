package main

import (
	"encoding/json"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"log/slog"
	"os"
	"sync"
	"time"
)

var ilog = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

func main() {
	globalConfig := NewOptions().ParseFlag()
	bytes, _ := json.MarshalIndent(globalConfig, "", "  ")
	ilog.Println("load config:\n" + string(bytes))

	if globalConfig.HttpLog == false {
		ilog.Print("No http logging enabled")
		os.Exit(0)
	}

	rolling := &lumberjack.Logger{
		Filename:   globalConfig.FileName,
		MaxSize:    20,
		MaxAge:     3,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   false,
	}

	//jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	jsonHandler := slog.NewJSONHandler(rolling, nil)
	jsonLogger := slog.New(jsonHandler)

	var wg sync.WaitGroup

	if globalConfig.HttpGetLog {
		wg.Add(1)
		// send get request
		go func() {
			defer wg.Done()
			for {
				time.Sleep(500 * time.Millisecond)
				printGetLog(jsonLogger)
			}
		}()
	}

	if globalConfig.HttpPostLog {
		wg.Add(1)
		// send post request
		go func() {
			defer wg.Done()
			for {
				time.Sleep(500 * time.Millisecond)
				printPostLog(jsonLogger)
			}
		}()
	}

	if globalConfig.HttpPutLog {
		wg.Add(1)
		// send put request
		go func() {
			defer wg.Done()
			for {
				time.Sleep(500 * time.Millisecond)
				printPutLog(jsonLogger)
			}
		}()
	}
	if globalConfig.HttpDeleteLog {
		wg.Add(1)
		// send delete request
		go func() {
			defer wg.Done()
			for {
				time.Sleep(500 * time.Millisecond)
				printDeleteLog(jsonLogger)
			}
		}()
	}

	wg.Wait()

}
