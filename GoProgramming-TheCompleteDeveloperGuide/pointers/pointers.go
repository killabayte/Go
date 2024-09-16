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
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {
	shirt := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	backpack := Item{"Backpack", Active}
	watch := Item{"Watch", Active}

	items := []Item{shirt, pants, backpack, watch}
	fmt.Println(items)

	deactivate(&items[0].tag)
}
