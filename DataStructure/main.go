package main

import (
	"DataStructure/SortAlgorithm"
	"fmt"
)

func main() {
	// arr := []int{4, 3, 3, 5, 1, 4, 4, 4, 3}
	// for k, v := range arr { //可以只接收key
	// 	fmt.Println(k, v)
	// }
	// SortAlgorithm.SelectSort(arr)
	// SortAlgorithm.InsertSort(arr)
	// 	SortAlgorithm.BubbleSort(arr)
	// var arr []int
	// for i := 0; i < 100; i++ {
	// 	arr = append(arr, rand.Intn(100))
	// }
	// SortAlgorithm.SonicSort(arr)
	// SortAlgorithm.OESort(arr)
	// SortAlgorithm.HeapSort(arr)
	// fmt.Println(arr)
	// mp := map[string]int{ //无序排列，无法进行计数且按次序排序(如果出现相等的成绩)
	// 	"小明": 90,
	// 	"小白": 94,
	// 	"小李": 95,
	// 	"小黄": 99,
	// 	"小黑": 94,
	// }
	db1 := []string{
		"小明", "小白", "小李", "小黄", "小黑", "小红", "小慧", "小灰",
	}
	db2 := []int{
		90, 94, 95, 99, 95, 98, 95, 90,
	}
	res := SortAlgorithm.CountSort(db1, db2)
	fmt.Println(res, len(res))
}
