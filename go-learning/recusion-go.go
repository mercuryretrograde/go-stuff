package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	factoralSeed := 5
	fibonacciSeed := 8 // This value is the number in the sequence
	fmt.Println("Factoral", factoralSeed, "=", fact(factoralSeed))
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fib(n-1) + fib(n-2)
		}
	}
	fmt.Println("Fibonacci", fibonacciSeed, "=", fib(fibonacciSeed))

}
