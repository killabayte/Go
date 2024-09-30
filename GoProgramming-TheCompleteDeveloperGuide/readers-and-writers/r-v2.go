package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 5)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

}
