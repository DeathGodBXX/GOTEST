package main

import "fmt"

func main() {
	fmt.Printf("start\n")
	defer fmt.Printf("1\n")
	defer fmt.Printf("2\n")
	defer fmt.Printf("3\n")
	fmt.Printf("end\n")

	// defer执行效果:
	//1.返回值赋值
	//2.defer执行
	//3.RET指令

	//defer执行场景:文件句柄，数据库链接，网络链接
	fmt.Printf("\n======defer使用=========\n")
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //6
	fmt.Println(f4()) //5
	fmt.Println(f5()) //5
	fmt.Printf("\n==========使用结束===========\n")

}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	x = 5
	defer func() {
		x++
	}()
	return x
}

func f3() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f4() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f5() (x int) {
	x = 5
	defer func(x int) {
		x++
	}(x)
	return x
}
