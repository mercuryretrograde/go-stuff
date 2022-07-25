package main

import (
	"fmt"
	fp "path/filepath"
	"strings"
)

func main() {

	p := fp.Join("dir1", "dir2", "filename")
	fmt.Println("This is the file path:", p)

	fmt.Println(fp.Join("dir//", "filename"))
	fmt.Println(fp.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", fp.Dir(p))
	fmt.Println("Base(p):", fp.Base(p))

	fmt.Println(fp.IsAbs("dir/file"))
	fmt.Println(fp.IsAbs("/dir/file"))

	filename := "config.json"

	ext := fp.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := fp.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = fp.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

}
