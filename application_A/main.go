package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Openfile("aardvark.txt", os.O_RDONLY, os.FileMode(0600))
	check(err)
	defer file.Close()
	scanner := bufio.NewScaner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())
}
