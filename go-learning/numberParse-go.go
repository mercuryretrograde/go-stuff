package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println()

	f, _ := strconv.ParseFloat("1.245", 64)
	fmt.Println("Parsed string to float:", f)

	i, _ := strconv.ParseInt("1253", 0, 64)
	fmt.Println("Parsed string to int:", i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println("Parsed hex string to int:", d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println("Parsed string to unsigned int:", u)

	k, _ := strconv.Atoi("3022")
	fmt.Println("using strconv.Atoi for basic base 10 int parsing:", k)

	_, e := strconv.Atoi("wat")
	fmt.Println("Error after unsuccessful parse:", e)

	fmt.Println()

}
