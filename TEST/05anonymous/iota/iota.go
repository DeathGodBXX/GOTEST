// package iota //如果不算package main,package iota的main函数被忽略了吗？
package main

import "fmt"

func main() {
	const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
	const (
		t1, t2 = iota + 1, iota + 2 //每一行结尾添加分号用于区分行与行之间内容
		t3, t4 = iota + 1, iota + 5
		// 1，2，2，6
	)
	const t5 = iota //0
	// var (
	// 	e1 = 1
	// 	e2
	// 	e3
	// ) //const同时声明多个常量时，如果省略了值则表示和上面一行的值相同,对于var 不适用。。

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	fmt.Printf("\t%d\t%d\t%d\t%d\n", n1, n2, n3, n4)
	fmt.Printf("\t%d\t%d\t%d\t%d\t%d\n", t1, t2, t3, t4, t5)
	fmt.Printf("\n%b\n%b\n%b\n%b\n%b\n", KB, MB, GB, TB, PB)
	fmt.Printf("\n%d\n%d\n%d\n%d\n%d\n", KB, MB, GB, TB, PB)

}
