package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	fmt.Print("i is assigned ", i, " at the start and in text is: ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a week day")
	}

	tn := time.Now()
	switch {
	case tn.Hour() < 12:
		fmt.Println("It is the morning")
	default:
		fmt.Println("It is the afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
