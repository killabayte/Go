package main

import (
	"fmt"
	"os"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}
func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}
