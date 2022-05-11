package main

import "fmt"

func plustwo(a, b int) int {
	return a + b
}

func returnthree(c, d, e int) (int, int, int) {
	return c, d, e
}

func main() {
	fmt.Println("hello functions")
	for kounter := 0; kounter <= 5; kounter++ {
		zz := kounter
		result := plustwo(zz, 10)
		fmt.Printf("This is the kounter value %v\n", kounter)
		fmt.Printf("This is the function result %v\n", result)
	}
	x, y, z := returnthree(1, 2, 3)
	fmt.Println(x, y, z)
}
