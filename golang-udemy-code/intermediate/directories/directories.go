package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}

}

func main() {

	// err := os.Mkdir("directories/subdir", 0755)
	// checkError(err)

	// checkError(os.Mkdir("directories/subdir1", 0755))
	// defer os.RemoveAll("directories/subdir1")

	// checkError(os.WriteFile("directories/subdir1/file.txt", []byte("Hello into File"), 0755))

	// checkError(os.MkdirAll("directories/subdir1/child1", 0755))
	// checkError(os.MkdirAll("directories/subdir1/child2", 0755))
	// checkError(os.MkdirAll("directories/subdir1/child3", 0755))
	// checkError(os.MkdirAll("directories/subdir1/child4", 0755))
	// checkError(os.MkdirAll("directories/subdir1/child5", 0755))

	// os.WriteFile("directories/subdir1/file", []byte ("Hello from Directories"), 0755)
	// os.WriteFile("directories/subdir1/child1/file", []byte ("Hello from\n directories"), 0755)


	result, err := os.ReadDir("directories/subdir1")
	checkError(err)
	
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.Type(), entry.IsDir())
	}

	checkError(os.Chdir("directories/subdir1/child1"))

	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading directories/subdir/child1")
	for _, entry := range result {
		fmt.Println(entry)
	}

	os.Chdir("../../")

	dir, err := os.Getwd()
	checkError(err)

	fmt.Println(dir)

	// filepath.Walk and filepath.WalkDir
	fmt.Println("walking directory")
	pathfile := "subdir1"

	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	checkError(err)

	checkError(os.RemoveAll("subdir1"))


}
