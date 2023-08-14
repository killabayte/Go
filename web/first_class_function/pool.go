package main

import "fmt"

func callFunction(passedFunction func()) {
	passedFunction()
}
func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}
func callWithArguments(passedFunction func(string, bool)) {
	passedFunction("This sentence is", false)
}
func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}
