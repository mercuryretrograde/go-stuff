package main

import (
	"fmt"
)

func intVal(ival int) {
	ival = 0
}

func intPtr(iptr *int) {
	*iptr = 0
}

func setIntPtr(iptr *int) {
	*iptr = 50
}

func main() {
	seed := 10
	fmt.Println("initial value", seed)
	intVal(seed)
	fmt.Println("seed after value function call", seed)
	intPtr(&seed)
	fmt.Println("seed after pointer function call", seed)
	setIntPtr(&seed)
	fmt.Println("seed after set pointer function call", seed)
	fmt.Println("seed pointer value", &seed)
}
