package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	common := make(chan struct{})
	start := time.Now()
	var count int

	// ping
	go func() {
		// init the ping
		common <- struct{}{}
		for {
			count++
			common <- <-common
		}
	}()

	// pong
	go func() {
		for {
			common <- <-common
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("Rte: %.2f requests per seconds\n", float64(count)/time.Since(start).Seconds())
}
