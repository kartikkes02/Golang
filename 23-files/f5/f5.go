// read from one file to write into the another file

package main

import (
	"bufio"
	"os"
)

func main() {
	sourceFile, err := os.Open("f5i.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create("f5ii.txt")
	if err != nil {
		panic(err)
	}
	defer destinationFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destinationFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}
	writer.Flush() // Flushes any buffered data to the underlying writer

	// deleting a file

	// er := os.Remove("f5ii.txt")
	// if er != nil {
	// 	panic(er)
	// }
}
