// package main

// type S struct {
// 	M *int
// }

// func main() {
// 	var x S
// 	var i int
// 	ref(&i, &x)
// }
// func ref(y *int, z *S) { z.M = y }
// 编译器不知道y和z的关系，被y分配到堆上，z分配到栈上

package main

import "fmt"

type User struct {
	name interface{}
}

func main() {
	name := "WilburXu"
	// fmt.Printf("%p\n", &name)
	MyPrintln(name)
}

func MyPrintln(one interface{}) (n int, err error) {
	fmt.Printf("%p\n，%#v\n", &one, one) //值传递
	var userInfo = new(User)
	userInfo.name = one // 泛型赋值 逃逸咯
	// fmt.Printf("%p\n", &(userInfo.name)) //值传递
	return
}

// # command-line-arguments
// ./main.go:28:16: moved to heap: one
// ./main.go:29:12: ... argument does not escape
// ./main.go:30:20: new(User) does not escape
// ./main.go:25:11: name escapes to heap
// 0xc0000961f0
// ，"WilburXu"
// name虽然没有外部引用，遍历，但是对于interface类型，很遗憾，编译器在编译的时候很难知道在函数的调用或者结构体的赋值过程会是怎么类型，因此只能分配到堆上。

// package main

// type User struct {
// 	name string
// }

// func main() {
// 	name := "WilburXu"
// 	MyPrintln(name)
// }

// func MyPrintln(one string) (n int, err error) {
// 	var userInfo = new(User)
// 	userInfo.name = one // 泛型赋值 逃逸咯
// 	return
// }

// name虽然没有外部引用，遍历，但是对于interface类型，很遗憾，编译器在编译的时候很难知道在函数的调用或者结构体的赋值过程会是怎么类型，因此只能分配到堆上。当name,one改为string类型，确定化了，所以直接放在栈上了
