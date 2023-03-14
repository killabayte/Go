package main

import "fmt"

func main() {
	var arr1 [7]int = [7]int{7, 3, 6, 0, 4, 9, 10}
	fmt.Println(arr1)
	fmt.Printf("%d %d\n", len(arr1), cap(arr1))
	var arr2 []int = arr1[1:3]
	fmt.Println(arr2)
	fmt.Printf("%d %d\n", len(arr2), cap(arr2))
	arr2 = arr2[0 : len(arr2)+2]
	fmt.Println(arr2)
	fmt.Printf("%d %d\n", len(arr2), cap(arr2))
}
