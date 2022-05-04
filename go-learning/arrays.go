package main

import "fmt"

func main() {
	var myArray [5]int
	fmt.Println("emp:", myArray)

	myArray[4] = 100
	fmt.Println("set: ", myArray)
	fmt.Println("get: ", myArray[4])

	fmt.Println("len: ", len(myArray))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
