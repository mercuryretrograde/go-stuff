package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")
	sleepTime := time.Second
	fmt.Println("sleeping", sleepTime)
	time.Sleep(sleepTime)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done

}
