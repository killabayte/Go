package main

import "fmt"

type Counter struct {
	hits int
}

func increment(c *Counter) {
	c.hits++
	fmt.Println("Counter", c)
}

func replace(old *string, new string, c *Counter) {
	*old = new
	increment(c)
}

func main_example() {
	counter := Counter{}
	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)
}
