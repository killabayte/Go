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
	for k := range arr2 {
		arr2[k] += 1
	}
	fmt.Println(arr2)
	fmt.Printf("%d %d\n", len(arr2), cap(arr2))
	fmt.Println(arr1)

	var arr3 []int
	fmt.Println(arr3)
	fmt.Printf("%d %d\n", len(arr3), cap(arr3))
	arr3 = append(arr3, 5)
	fmt.Println(arr3)
	fmt.Printf("%d %d\n", len(arr3), cap(arr3))

	//Arrayarr1
	arr4 := make([]int, 3)
	fmt.Println(arr4)
	fmt.Printf("%d %d\n", len(arr4), cap(arr4))

	//Array2
	arr5 := make([]int, 11)
	fmt.Println(arr5)
	fmt.Printf("%d %d\n", len(arr4), cap(arr5))

	//Conclusion: Slice are more popular structure in Go
	//But arrays can be used for optimize resources usage and so on...
}
