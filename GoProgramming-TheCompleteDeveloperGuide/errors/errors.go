package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("No element at index %v", index))
	} else {
		return s.values[index], nil
	}
}

func main() {

}
