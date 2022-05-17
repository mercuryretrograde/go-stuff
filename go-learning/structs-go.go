package main

import "fmt"

type person struct {
	fname string
	age   int
}

func newPerson(name string) *person {
	p := person{fname: name}
	p.age = 20
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{fname: "Alice", age: 30})

	fmt.Println(person{fname: "Fred"})

	fmt.Println(&person{fname: "Ann", age: 100})

	fmt.Println(newPerson("Jon"))

	s := person{fname: "Sean", age: 20}
	fmt.Println(s.fname)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}
