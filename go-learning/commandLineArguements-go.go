package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProgs := os.Args
	argsWithoutProg := os.Args[1:]
	args := os.Args[3]

	fmt.Println("full command line", argsWithProgs)
	fmt.Println("arguments without program name", argsWithoutProg)
	fmt.Println("third argument", args)

}
