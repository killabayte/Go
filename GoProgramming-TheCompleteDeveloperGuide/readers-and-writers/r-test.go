package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

command_count := 0
non_blanck_count := 0

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		input, inputErr := r.ReadString('\n')
		l := strings.TrimSpace(input)
		if inputErr == io.EOF {
			break
		}
		if len(l) > 0{
			non_blanck_count ++
		}
		switch l {
		case l == "hello":
			command_count ++
			fmt.Println("Hi User, nice to meet you!")
		case l == "bye":
			command_count ++
			fmt.Println("Have a good day!And come back soon")
		}
	}
}
