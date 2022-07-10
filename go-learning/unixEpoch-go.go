package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println("Now:", now)

	fmt.Println("Now Unix:", now.Unix())

	fmt.Println("Unix milli:", now.UnixMilli())
	fmt.Println("Unix nano", now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))

}
