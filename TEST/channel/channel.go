package main

import (
	"fmt"
	"time"
)

func main() {
	// var ch chan []int
	ch1 := make(chan string, 4)
	ch1 <- "str"
	x := <-ch1
	close(ch1)
	fmt.Printf("%v\n", x)

	//无缓冲通道，主线程发送，必须先执行接收操作，否则不知道接收方，无法发出
	ch := make(chan int)
	go func() {
		x := <-ch
		fmt.Printf("接收完成，%v\n", x)
	}()
	ch <- 10
	fmt.Println("发送成功")
	time.Sleep(time.Second)

	//无缓冲通道，主线程接收，必须先知道发送方，方可接收
	ch_recv := make(chan int)
	go func() {
		ch_recv <- 10
		fmt.Println("发送完成")
	}()
	z := <-ch_recv
	fmt.Println("接收成功，", z)

	//无缓冲通道，开启两个goroutine，发送接收处理，再发送接收操作
	ch3 := make(chan int)
	ch4 := make(chan int)
	go func() {
		for i := 20; i >= 0; i-- {
			ch3 <- i
		}
		close(ch3)
	}()
	go func() {
		for i := range ch3 { //for range取值,通道关闭前的数据都能读取接收
			ch4 <- i * i
		}
		close(ch4)
	}()
	for {
		i, ok := <-ch4 //等价于for range取值，读取通道关闭前的数据
		if !ok {
			break
		}
		fmt.Printf("%v\t", i)
	}
}
