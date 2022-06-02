package main

import (
	"fmt"
	"time"
)

var chan_one = make(chan string)
var chan_two = make(chan string)

func f_one() {
	time.Sleep(1 * time.Second)
	chan_one <- "one"
}

func f_two() {
	time.Sleep(2 * time.Second)
	chan_two <- "two"

}

func main() {
	go f_one()

	go f_two()

	for idx := 0; idx < 2; idx++ {
		select {
		case message_one := <-chan_one:
			fmt.Println("received", message_one)
		case message_two := <-chan_two:
			fmt.Println("received", message_two)
		}
	}
}
