package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	command_count := 0
	non_blanck_count := 0

	r := bufio.NewReader(os.Stdin)

	for {
		input, inputErr := r.ReadString('\n')
		l := strings.TrimSpace(input)
		if inputErr == io.EOF {
			break
		}
		if len(l) > 0 {
			non_blanck_count++
		}
		switch l {
		case "hello":
			command_count++
			fmt.Println("Hi User, nice to meet you!")
		case "bye":
			command_count++
			fmt.Println("Have a good day!And come back soon")
		}
	}
	fmt.Println("The number of commands entered:", command_count)
	fmt.Println("The number of non-blank lines:", non_blanck_count)
}
