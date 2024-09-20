package main

import "fmt"

type Calculator byte

const (
	Add Calculator = iota
	Subtract
	Multiply
	Divide
)

func (c Calculator) Calculate(n1, n2 int) int {
	switch c {
	case Add:
		return (n1 + n2)
	case Subtract:
		return (n1 - n2)
	case Multiply:
		return (n1 * n2)
	case Divide:
		return (n1 / n2)
	default:
		return 0
	}
}

func main() {
	addOperation := Add.Calculate(42, 42)
	fmt.Println(addOperation)
	subOperation := Subtract.Calculate(addOperation, 42)
	fmt.Println(subOperation)

}
