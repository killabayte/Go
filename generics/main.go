package main

import (
	"fmt"
	"reflect"
)

func main() {
	var t1 int = 123
	fmt.Printf("plusOne: %v (type: %s)\n", plusOne(t1), reflect.TypeOf(plusOne(t1)))
	var t2 float64 = 123.12
	fmt.Printf("plusOne: %v (type: %s)\n", plusOne(t2), reflect.TypeOf(plusOne(t2)))
}

func plusOne[V int | float64 | int64 | float32 | int32](t V) V {
	return t + 1
}
