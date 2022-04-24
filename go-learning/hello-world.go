package main

import "fmt"

func main() {
	fmt.Println("hello world")
	for kounter := 0; kounter <= 5; kounter++ {
		fmt.Printf("This is the kounter value %v\n", kounter)
	}
	fmt.Println("bye world")

}
