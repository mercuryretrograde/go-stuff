package main

import (
	"fmt"
)

type minType interface {
	int32 | int64 | string
}

func chooseMin[R minType](x, y R) R {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	var valA string = "a"
	var valB string = "t"
	fmt.Println("generic test code start")
	fmt.Println("generic test result", chooseMin(valA, valB))
}
