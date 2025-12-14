package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var shutdownHook []context.CancelFunc
	for _, apiLogMock := range ApiLogMockSlice {
		worker := NewLogWorker(apiLogMock)
		cancelFunc := worker.Run()
		shutdownHook = append(shutdownHook, cancelFunc)
	}

	// 监听系统信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	fmt.Println("execute shutdown hook")
	for _, hook := range shutdownHook {
		hook()
	}
	fmt.Println("wait shutdown hook")
	time.Sleep(1 * time.Second)

}

type LogWorker struct {
	mock    *ApiLogMock
	running bool
	lock    sync.Mutex
}

func NewLogWorker(mock ApiLogMock) *LogWorker {
	return &LogWorker{
		mock: &mock,
	}
}

func (w *LogWorker) Run() context.CancelFunc {
	if w.running {
		panic("already running")
	}
	for i := 0; i < 5; i++ {
		if w.lock.TryLock() {
			ctx, cancelFunc := context.WithCancel(context.Background())
			go w.doTask(ctx)
			w.running = true
			w.lock.Unlock()
			return cancelFunc
		}
		if w.running {
			panic("already running")
		}
	}
	return nil
}

func (w *LogWorker) doTask(ctx context.Context) {
	endpoint := w.mock.Endpoint
	url := w.mock.Url
	status := w.mock.Status
	qps := w.mock.Qps

	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			return
		default:
		}
		// task
		for range qps {
			Logger.Info("",
				"endpoint", endpoint,
				"url", url,
				"status", status,
				"costTime", CostTimeRand(),
			)
		}
		time.Sleep(time.Second * time.Duration(1))
	}
}
