package main

func main() {
	// Do nothing
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}
