package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	myString := "some long thing"

	p("Contains:", s.Contains(myString, "me"))
	p("Count:", s.Count(myString, "o"))

	// p("HasPrefix: ", s.HasPrefix("test", "te"))
	// p("HasSuffix: ", s.HasSuffix("test", "st"))
	// p("Index:     ", s.Index("test", "e"))
	// p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	// p("Repeat:    ", s.Repeat("a", 5))
	// p("Replace:   ", s.Replace("foo", "o", "0", -1))
	// p("Replace:   ", s.Replace("foo", "o", "0", 1))
	// p("Split:     ", s.Split("a-b-c-d-e", "-"))
	// p("ToLower:   ", s.ToLower("TEST"))
	// p("ToUpper:   ", s.ToUpper("test"))
}
