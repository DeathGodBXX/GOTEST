package main

import (
	"DataStructure/SortAlgorithm"
	"fmt"
	"math/rand"
)

func main() {
	// arr := []int{4, 3, 3, 5, 1, 4, 4, 4, 3}
	// SortAlgorithm.SelectSort(arr)
	// SortAlgorithm.InsertSort(arr)
	// 	SortAlgorithm.BubbleSort(arr)
	var arr []int
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}
	// SortAlgorithm.SonicSort(arr)
	// SortAlgorithm.OESort(arr)
	SortAlgorithm.HeapSort(arr)
	fmt.Println(arr)
}
