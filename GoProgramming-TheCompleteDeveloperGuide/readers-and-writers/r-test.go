package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func quit(nbc, cc int) {
	fmt.Println("Gracefully shutting down...")
	fmt.Printf("Number of non-blank lines: %d\n", nbc)
	fmt.Println("The number of commands entered:", cc)
}

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
		case "Q":
			quit(non_blanck_count, command_count)
			return
		case "q":
			quit(non_blanck_count, command_count)
			return
		}
	}
}
