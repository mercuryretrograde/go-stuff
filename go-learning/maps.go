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
		myMap[keyName] = i
	}
	fmt.Println(myMap)
	delete(myMap, "k5")
	fmt.Println(myMap)
	oneItem, _ := myMap["k4"]
	fmt.Println(oneItem)

}
