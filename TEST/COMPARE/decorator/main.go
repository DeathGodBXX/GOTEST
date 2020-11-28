package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	dec := decorator(add)
	count := 100000
	for i := 0; i < count; i++ {
		dec(1, i)
	}
	collapsed := time.Since(start) / 1e6
	fmt.Printf("总耗时:%v \n", collapsed)
}

func decorator(a func(int, int) int) func(int, int) {
	wrapper := func(x, y int) {
		fmt.Printf("\n=====程序开始======\n")
		fmt.Printf("%d+%d=%d\n", x, y, a(x, y))
		fmt.Printf("\n=====程序结束======\n")
	}
	return wrapper
}

func add(x, y int) int {
	return x + y
}
