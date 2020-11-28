package main

import (
	"fmt"
	"time"
)

func main() {
	N := 100
	start := time.Now().UnixNano()
	forUsage(N)
	collapse := time.Now().UnixNano() - start

	start1 := time.Now().UnixNano()
	forUsage2(N)
	collapse1 := time.Now().UnixNano() - start1

	start2 := time.Now().UnixNano()
	var x []int
	forUsage3(x, N)
	collapse2 := time.Now().UnixNano() - start2

	fmt.Printf("b in for collapsed:%#v ns,\nb1 in for collapsed:%#v ns\n,collased:%#v\n", collapse, collapse1, collapse2)
	// alt+z 续行操作
}
func forUsage(count int) {
	for i := 0; i < count; i++ {
		b := []int{1, 2, 3, 4, 5, 6, 7, 8} //不逃逸
		_ = b
	}
}

func forUsage2(count int) {
	var b1 []int
	for i := 0; i < count; i++ {
		b1 = []int{1, 2, 3, 4, 5, 6, 7, 8} // 逃逸
		_ = b1
	}
}

func forUsage3(x []int, N int) { //不逃逸
	for i := 0; i < N; i++ {
		x = []int{1, 2, 3, 4, 5, 6, 7, 8} //逃逸
		_ = x                             //最终x的数据还是会逃逸的
	}
}

// ./main.go:28:13: []int literal does not escape
// ./main.go:36:13: []int literal escapes to heap
// ./main.go:41:16: x does not escape
// ./main.go:43:12: []int literal escapes to heap
// ./main.go:23:12: ... argument does not escape
// ./main.go:23:13: collapse escapes to heap
// ./main.go:23:13: collapse1 escapes to heap
// ./main.go:23:13: collapse2 escapes to heap
