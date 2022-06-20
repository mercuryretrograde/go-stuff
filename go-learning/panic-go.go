package main

import (
	"os"
)

func main() {

	//panic("a problem")

	_, err := os.Create("/tmp/file.dat")
	if err != nil {
		panic(err)
	}
}
