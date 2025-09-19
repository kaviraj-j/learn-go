package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("app started")
	myCh := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("after timeout")
		myCh <- false
	}()
	fmt.Println("app ended")
	<-myCh
}
