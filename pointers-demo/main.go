package main

import "fmt"

func main() {
	a := []string{"string"}
	testPointer(a)
	fmt.Printf("a: %v\n", a)
}

func testPointer(a []string) {
	a[0] = "changed element"
	a = append(a, "another string")
}
