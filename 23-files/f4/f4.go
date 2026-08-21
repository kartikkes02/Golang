package main

import (
	"os"
)

func main() {
	f, err := os.Create("f4.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write into the files

	// f.WriteString("Hello Go")
	// f.WriteString(" ,Goalng")

	// replace
	bytes := []byte("He")
	f.Write(bytes) // WriteAt writes len(b) bytes to the File starting at byte offset off.
}
