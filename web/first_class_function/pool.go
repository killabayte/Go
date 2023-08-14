package main

func callFunction(passedFunction func()) {
	passedFunction()
}
func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}
