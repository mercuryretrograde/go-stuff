package main

import (
	"fmt"
)

func main() {
	mySlice := make([]string, 3)
	fmt.Println("my slice is: ", mySlice)
	mySlice[0] = "hello"
	mySlice[1] = "thought process"
	mySlice[2] = "a"
	fmt.Println("with values assigned = ", mySlice)
	mySlice = append(mySlice, "new string")
	mySlice = append(mySlice, "easy as")
	fmt.Println(mySlice)
	littleBit := mySlice[1:3]
	fmt.Print(littleBit)

}
