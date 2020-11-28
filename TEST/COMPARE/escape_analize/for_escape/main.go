package main

import (
	"fmt"
	"time"
)

func main() {
	N := 100
	start := time.Now().UnixNano()
	for i := 0; i < N; i++ {
		b := []int{1, 2, 3, 4, 5, 6, 7, 8}
		_ = b
	}
	collapse := time.Now().UnixNano() - start

	start1 := time.Now().UnixNano()
	var b1 []int
	for i := 0; i < N; i++ {
		b1 = []int{1, 2, 3, 4, 5, 6, 7, 8}
		_ = b1
	}
	collapse1 := time.Now().UnixNano() - start1

	start2 := time.Now().UnixNano()
	var x []int
	forUsage(x, N)
	collapse2 := time.Now().UnixNano() - start2
	fmt.Printf("b in for collapsed:%#v ns,\nb1 in for collapsed:%#v ns\n,collased:%#v\n", collapse, collapse1, collapse2)
	// alt+z 续行操作
	// x = []int{1, 2, 3, 4}
	// fmt.Printf("%p\n", x)
}

func forUsage(x []int, N int) {
	for i := 0; i < N; i++ {
		x = []int{1, 2, 3, 4, 5, 6, 7, 8}
		_ = x
	}
	// fmt.Printf("%p\n", x)
	// fmt.Printf("%#v\n", x)
}
