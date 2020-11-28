package main

import (
	"fmt"
	"time"
)

//N is constant
const N = 3

func main() {
	start := time.Now()
	for i := 0; i < N; i++ {
		test()
	}
	elapsed := time.Since(start) / N
	fmt.Println("总用时:", elapsed)
}

func test() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d = %d  ", i, j, i*j)
		}
		fmt.Println()
	}

}
