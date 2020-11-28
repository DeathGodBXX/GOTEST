package main

import (
	"fmt"
	"math"
)

func main1() {
	a := []int{1, 2, 3, 4}
	// Sort(a[1:])
	p := &a
	Sort1(p)
	fmt.Printf("In main,a=%#v\n", a) //引用类型

}

// func Sort(a []int) {
// 	if len(a) == 0 {
// 		fmt.Printf("a in nil\n")
// 		return
// 	}
// 	a[0] = 12345
// 	fmt.Printf("In Sort,a=%#v\n", a)
// }

func Sort1(p *[]int) {
	(*p)[1] = 123
	(*p)[2] = 289
}

func main() {
	l, r := 10.8, 19.4
	s := int(math.Floor((float64(l+r) / 2)))
	fmt.Println(s)
}
