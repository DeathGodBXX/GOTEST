package main

import "fmt"

func main() {
	slice := make([]interface{}, 1, 10)
	//make申明的元素，如果使用[]设定值，必须保证size大于等于保存之后的size,否则越界；而append会扩容
	slice[0] = 10
	// slice[1] = 11
	slice = append(slice, 12)
	fmt.Println(slice)
}
