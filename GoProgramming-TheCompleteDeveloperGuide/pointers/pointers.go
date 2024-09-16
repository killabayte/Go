package main

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool
type Item struct {
	name string
	tag  SecurityTag
}
