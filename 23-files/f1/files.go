package main

import ( 
	"fmt"
	"os"
)

func main() {
	fmt.Println("Files")
	f, err := os.Open("files1.txt")

	if err != nil {
		panic(err)
		// panic is a built-in function that stops the normal execution of a program and begins panicking.
		// When the function is called, it prints a message to the console and then stops the program.
	}

	// f.Stat() returns a FileInfo object that contains information about the file, such as its size,
	// name, and permissions.

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	// fileInfo provides information about the file, such as its size, name, and permissions.
	fmt.Println("File Size: ", fileInfo.Size())
	fmt.Println("File Name: ", fileInfo.Name())
	fmt.Println("File Mode: ", fileInfo.Mode())
	fmt.Println("File ModTime: ", fileInfo.ModTime())
	fmt.Println("File IsDir: ", fileInfo.IsDir())
}
