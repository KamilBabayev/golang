	package main
// Defer functions is for executing code later after some event happened.

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("names.txt")
	defer closeFile(file)   // defer func here runs after writing to file function in below line.
	writeToFile(file, "salam\n")
}

func createFile(path string) *os.File {
	fmt.Println("creating file")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func writeToFile(file *os.File, name string) {
	fmt.Println("Writing to file")
	fmt.Println(file, name)
	fmt.Println("wrote")
}

func closeFile(file *os.File) {
	fmt.Println("closeing file")
	file.Close()
}

