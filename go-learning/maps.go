package main

import (
	"fmt"
)

func main() {

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["k0"] = 0
	fmt.Println(myMap)

	for i := 1; i <= 10; i++ {
		var keyName string
		keyName = fmt.Sprint("k", i)
		fmt.Println(keyName)
		myMap[keyName] = i
	}
	fmt.Println(myMap)

}
