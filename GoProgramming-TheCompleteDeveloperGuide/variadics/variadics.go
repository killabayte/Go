package main

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	all := append(a, b...)
}
