package main

import (
	"fmt"
	"sort"
)

func main() {

	unsortedString := []string{"z", "a", "e", "t", "y"}
	sort.Strings(unsortedString)
	fmt.Println("Sorted strings:", unsortedString)

	unsortedIntegers := []int{3, 7, 1, 8, 0, 34, 678, -1}
	testOne := sort.IntsAreSorted(unsortedIntegers)
	fmt.Println("Integers are sorted =", testOne)

	sort.Ints(unsortedIntegers)
	fmt.Println("Sorted integers:", unsortedIntegers)
	testTwo := sort.IntsAreSorted(unsortedIntegers)
	fmt.Println("Integers are sorted =", testTwo)

}
