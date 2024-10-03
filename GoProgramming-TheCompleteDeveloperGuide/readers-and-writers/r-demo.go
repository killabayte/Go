package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
	}
}
