package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	someString := "abc123!?$*&()'-=@~"

	fmt.Println()
	fmt.Println("my string:", someString)

	sEnc := b64.StdEncoding.EncodeToString([]byte(someString))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(someString))
	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
