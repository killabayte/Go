package main

import "fmt"

type Part string

func showLine(line []Part) {
	for i := 0; i <= len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}
