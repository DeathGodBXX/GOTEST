package main

import (
	"fmt"
)

var (
	t1 = 1
	t2 string
	t3 int
	t4 byte
) // var 声明，不需要，分割开

const (
	pi=3.141592653
	e =2.5
)

func main() {
	x, _ := foo() // _匿名变量，如果要使用z替代_,go语言规则必须使用，但是有不想使用
	_, y := foo()
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
}

func foo() (int, string) { //后面的int,string用来说明返回值的类型
	return 10, "OK!"
}

//
//函数外的每个语句都必须以关键字开始（var、const、func等）
//:=不能使用在函数外。
//_多用于占位，表示忽略值。
