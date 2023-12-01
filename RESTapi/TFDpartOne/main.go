package main

func main() {
	// Do nothing
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans
}
