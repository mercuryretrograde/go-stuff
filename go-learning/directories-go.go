package main

import (
	"fmt"
	"os"
	fp "path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	directoryName := "subdir"
	err := os.Mkdir(directoryName, 0755)
	check(err)

	defer os.RemoveAll(directoryName)

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}	

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdirectory")
	err = fp.Walk("subdir", visit)

}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
