package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithRetirnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("MethodWithParameters called with", f)
}

func (m MyType) MethodWithRetirnValue() string {
	return "Hi from MethodWithRetirnValue"
}

func (my MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
