package main

import (
	"fmt"
)

func vari_sum(nums ...int) int {
	total := 0
	fmt.Println("nums = ", nums)
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	result := vari_sum(4, 5, 6, 7, 8)
	fmt.Println("sum returned =", result)
}
