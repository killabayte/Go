package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "John Week"
	employee.Salary = 999999
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}
