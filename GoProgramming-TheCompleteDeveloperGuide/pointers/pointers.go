package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool
type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}
func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("Checking out")
}
