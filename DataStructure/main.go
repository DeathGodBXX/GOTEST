package main

import (
	"DataStructure/SortAlgorithm"
	"fmt"
)

func main() {
	arr := []int{5, 7, 3, 2, 1, 10, 4, 6, 8, 9}
	// SortAlgorithm.SelectSort(arr)
	// fmt.Println(arr)
	// SortAlgorithm.InsertSort(arr)
	// fmt.Println(arr)
	SortAlgorithm.BubbleSort(arr)
	fmt.Println(arr)
}
