package main

import "fmt"

func main() {
	FuncA()
	FuncB()
	FuncC()

}

//FuncA func
func FuncA() {
	fmt.Printf("Hello,I'm FuncA()\n")
}

//FuncB func
func FuncB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("recover in B\n")
		}

	}()
	panic("something is wrong")
	fmt.Printf("Hello,I'm FuncB()\n")
}

//FuncC func
func FuncC() {
	fmt.Printf("Hello,I'm FuncC()\n")
}

//recover必须搭配defer使用
//defer必须在panic之后
