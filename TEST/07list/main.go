package main

import "fmt"

func main() {
	//1.定义变量并初始化
	//数组长度属于数组的一部分
	var a [3]int
	// var aa [3]string
	var b = [3]int{1, 5}
	var c = [3]string{"北京", "上海", "深圳"}
	d := "hello"
	//2.根据初始值个数动态推断数组长度 半动态性 长度推断 一旦确定,不可更改 outof range
	var e = [...]string{"北京", "上海", "深圳"}
	var f = [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
	fmt.Printf("%T\n%T\n", e, f)
	//3.以索引的方式初始化数组
	var g = [...]int{1: 1, 3: 3}
	g1 := [...]string{1: "ok", 3: "yes"}
	g2 := [...]string{} //声明不定长字符串,以空初始化字符串
	// g2[1] = "happy"
	var g3 [4]string
	// var g4 [...]string //不可行
	//要么一开始给定数组长度,要么利用初始化,自行推断数组长度(不能越界)
	fmt.Println(g, g1, g2, g3)
	fmt.Printf("%T\t%T\n", g, g1)

	//遍历数组c
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
	for _, value := range c {
		fmt.Println(value)
	}

	//多维数组,只有第一层可以省略
	a1 := [...][2]string{
		{"北京", "上海"},
		{"广东", "深圳"},
		{"佛山", "陕西"},
	}
	fmt.Println(a1, a1[2][1])
	for _, v1 := range a1 {
		for _, v2 := range v1 {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}
}
