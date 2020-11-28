package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

//defer执行效果:
//先把实参全部带进去，先执行内部函数，再延迟最外一层函数
//注意使用引用的场景，有可能后修改，能达到使用最后函数的效果
//1.x=1
//2.y=2
//3.calc("A",1,2)  "A" 1 2 3
//4.defer calc("AA",1,3)
//5.calc("B",10,2) "B" 10 2 12
//6.defer calc("BB",10,12)
//7.y=20
//执行第六步结果 "BB" 10 12 22
//执行第四步结果 "AA"  1  3  4

//最终结果：
//"A" 1 2 3
//"B" 10 2 12
//"BB" 10 12 22
//"AA"  1 3 4
