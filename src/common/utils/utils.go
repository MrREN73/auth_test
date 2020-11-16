package utils

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Stop is a shortcut for service stopper handler.
type Stop = func()

func waitSignal() {
	var sigch = make(chan os.Signal, 1)

	signal.Notify(sigch, syscall.SIGTERM)
	signal.Notify(sigch, syscall.SIGINT)

	<-sigch
}

// GracefulStop calls passed handler and waits for passed duration.
// If timeout exceeded stops roughly.
func GracefulStop(td time.Duration, stops ...Stop) {
	waitSignal()

	ctx, cancel := context.WithTimeout(context.Background(), td)
	defer cancel()

	sch := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(len(stops))

	stopping := func(stop Stop) {
		stop()
		wg.Done()
	}

	for _, stop := range stops {
		go stopping(stop)
	}

	waiting := func() {
		wg.Wait()
		sch <- struct{}{}
	}
	go waiting()

	select {
	case <-sch:
		log.Println("stopped gracefully")
	case <-ctx.Done():
		log.Fatal("time is out")
	}
}
