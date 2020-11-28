package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarter(ch2, ch1)
	printer(ch2)
	// for i := range ch2 {
	// 	fmt.Println(i)
	// }
}

func counter(out chan<- int) {
	//goroutine中需要关闭out,只要发送接收端，有一处在main中，都不需要关闭channel
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarter(out chan<- int, in <-chan int) { //
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
