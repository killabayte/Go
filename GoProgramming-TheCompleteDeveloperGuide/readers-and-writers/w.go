package main

import (
	"bytes"
	"fmt"
)

func main_old_v4() {
	buffer := bytes.NewBufferString("")
	numBytes, err := buffer.WriteString("SAMPLE")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Wrote %v bytes: %c\n", numBytes, buffer)
	}
}
