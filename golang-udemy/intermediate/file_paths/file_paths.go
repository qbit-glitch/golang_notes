package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	fmt.Println()
	relativePath := "./data/out.txt"
	absolutePath := "/home/user/docs/file.txt"

	// Join Paths using filepath.join
	joinedPath := filepath.Join("home", "documents", "downloads", "file.txt")
	fmt.Println("Joined Path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)



	dir, file := filepath.Split("/home/usr/docs/file.txt")
	fmt.Println("File:", file)
	fmt.Println("Dir:", dir)
	fmt.Println(filepath.Base("/home/user/docs/"))

	fmt.Println("Is relativePath variable absolute:", filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath variable absolute:", filepath.IsAbs(absolutePath))



	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))


	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)


	rel, err = filepath.Rel("a/c/d", "e/f/d")
	if err!= nil {
		panic(err)
	}
	fmt.Println(rel)



	absPath, err := filepath.Abs(relativePath)
	if err!= nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Absolute Path:", absPath)
	}

}