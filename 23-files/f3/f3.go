package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// if negative, it will read all the files in the directory.
	fileInfo, err := dir.Readdir(-1)

	for _, fi := range fileInfo {
		fmt.Println(fi.Name(), fi.IsDir())
	}
}
