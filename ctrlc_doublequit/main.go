package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	var interruptCount int32

	// Main loop ticker
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("app running...")

		case <-sigChan:
			// Increment interrupt count safely
			count := atomic.AddInt32(&interruptCount, 1)

			if count == 1 {
				fmt.Println("Ctrl+C pressed once, ignoring... (press again within 1s to quit)")

				// Reset after 1s
				go func() {
					time.Sleep(1 * time.Second)
					atomic.StoreInt32(&interruptCount, 0)
				}()

			} else {
				fmt.Println("Ctrl+C pressed twice, exiting...")
				os.Exit(0)
			}
		}
	}
}
