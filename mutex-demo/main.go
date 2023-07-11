package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type mytype struct {
	counter int
	mu      sync.Mutex
}

func main() {
	mytypeInstance := mytype{}
	finished := make(chan bool)
	for i := 0; i < 7; i++ {
		go func(mytypeInstane *mytype) {
			mytypeInstane.mu.Lock()
			fmt.Printf("Input Counter: %d\n", mytypeInstance.counter)
			mytypeInstane.counter++
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			if mytypeInstance.counter == 5 {
				fmt.Printf("Found counter equal: 5\n")
			}
			fmt.Printf("Output Counter: %d\n", mytypeInstance.counter)
			finished <- true
			mytypeInstane.mu.Unlock()
		}(&mytypeInstance)
	}
	for i := 0; i < 7; i++ {
		<-finished
	}
	fmt.Printf("Counter: %d\n", mytypeInstance.counter)
}
