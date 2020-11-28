package main

import "fmt"

func main() {
	// s := "Hello,everyone!中文"
	// for x, y := range s {
	// 	fmt.Printf("\t%d\t%c\t%T\n", x, y, y)
	// }
	// ss := []byte(s) //把字符串切割成byte,rune类型构成的列表或数组，
	// tt := []rune(s) //其中byte用于处理ascII码的，rune用于处理中文，日文，韩文等的utf8编码
	// fmt.Printf("%c\n%c\n", ss, tt)
	// fmt.Println(ss)
	New()
	Check()
	Finger()
}

//New method
func New() {
	if score := 80; score > 90 {
		fmt.Println("excellent")
	} else if score > 75 {
		fmt.Println("passed")
	} else {
		fmt.Println("learn again")
	}
}

//Check method
func Check() {
	s := "Hello"
	for i, v := range s {
		fmt.Printf("%d  %c \n", i, v)
	}
}

//Finger method
func Finger() {
	switch s := 3; s {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	}
}
