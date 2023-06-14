package main

import (
	"fmt"
	"reflect"
)

func main() {
	var t1 string = "This is a string"
	var t2 *string = &t1
	discoverType(t2)
	var t3 int = 42
	discoverType(t3)
	discoverType(nil)
}

func discoverType(t any) {
	switch v := t.(type) {
	case string:
		t2 := v + "..."
		fmt.Printf("String found: %s\n", t2)
	case *string:
		fmt.Printf("Pointer string found: %s\n", *v)
	case int:
		fmt.Printf("Int type found: %d\n", v)
	default:
		myType := reflect.TypeOf(t)
		if myType == nil {
			fmt.Printf("type is nil")
		} else {
			fmt.Printf("Type not found: %s\n", myType)
		}
	}
}
