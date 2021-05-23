package main

import (
	"Gin/example"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//必须要有主线程执行的一个gin,否则不会暴露出接口
	//这会使主线程阻塞等待客户端发送请求，否则执行完主线程直接结束，也就没有暴露接口，只能有一个goroutine
	//为何放在上面而不是下面，如果放在下面，那么主线程执行完直接结束了，串行执行，而go example.LoadHTML()和fmt.Println("")后者主线程执行完直接结束了
	if ex, err := os.Executable(); err == nil {
		fmt.Printf("filepath:%v", filepath.Dir(ex))
	}

	go example.Testexample()
	go example.DefinedSelfHTML()
	example.LoadHTML()

}
