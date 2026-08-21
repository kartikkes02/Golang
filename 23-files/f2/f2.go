package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("files2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close() // Close the file when the function returns

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	buf := make([]byte, fileInfo.Size()) // []byte. 12
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		fmt.Println(d, string(buf[i]))
	}

	// Another way to read a file
	// Using the os.ReadFile function,
	// which reads the entire file into memory and returns the data as a byte slice.
	// This is useful for small files, but can be inefficient for large files.

	// data, err := os.ReadFile("files2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))
}
