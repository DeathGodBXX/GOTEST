package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//1.simple use
	// go func() {
	// 	fmt.Println("hello")
	// }()
	// fmt.Printf("main thread\n")
	// time.Sleep(time.Second)

	//2.parallel excute
	for _, i := range "hello" {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}

func hello(i rune) {
	defer wg.Done()
	fmt.Printf("hello, %v(%c)\n", i, i)
}
