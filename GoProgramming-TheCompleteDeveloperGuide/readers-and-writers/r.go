package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main_old() {
	source := strings.NewReader("SAMPLE")
	buffered := bufio.NewReader(source)

	newString, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newString)
	} else {
		fmt.Println("Something went wrong")
	}
}
