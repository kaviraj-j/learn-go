package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	ticker := time.Tick(500 * time.Millisecond)
	quit := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		quit <- true
	}()
	for {
		select {
		case <-ticker:
			fmt.Println("app running...")
		case <-quit:
			fmt.Println("quitting...")
			time.Sleep(500 * time.Millisecond)
			os.Exit(0)
		}
	}
}
