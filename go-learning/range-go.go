package main

import (
	"fmt"
)

func main() {
	numList := []int{1, 2, 3, 4, 5, 15, 35, 21, 210}
	sum := 0
	for _, num := range numList {
		sum += num
	}
	fmt.Println("The sum of the range is ", sum)

	for _, num := range numList {
		if num%5 == 0 {
			fmt.Println("mod 5 for:", num)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cranberry"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Printf("key = %s \n", k)
	}

	for i, c := range "my name is che deane" {
		fmt.Println(i, c)
	}

}
