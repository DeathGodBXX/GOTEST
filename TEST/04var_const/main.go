package main

import "fmt"

var t int = 3 // 等价于var t = 3

var tt = 3 //全局变量可以不使用，但是局部变量必须要使用

// ttt := 5 not permitted outside of func
// 短变量声明，不能在全局变量中

func main() {
	var s int = 5
	var t = "I'm handsome!"
	ts := "Are you OK?"
	fmt.Println("HEONY!\n", s, "\n", t+"\n", ts+"\n") //add a newline
	fmt.Printf("\n%d\n%s\n%s\n", s, t, ts)            // use format to show results
	fmt.Print("\n", s, "\n", t, "\n", ts, "\n")       // not add a newline
}
