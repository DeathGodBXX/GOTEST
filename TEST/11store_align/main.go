package main

import (
	"fmt"
	"unsafe"
)

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

type Part2 struct {
	a bool
	b int8
	c byte
	d int32
	e int64
}

type Part3 struct {
	a bool
	b int8
	c byte
	d int64
	e int32
}

type test struct {
	number int
	data interface{}
}

type test1 struct{
	data interface{}
	number int
}

type stack struct{
	size int8
	data interface{}
	next *stack //指针也占用内存
}

func main() {
	part1 := Part1{}
	part2 := Part2{}
	part3 := Part3{}
	fmt.Printf("bool size:%d\n", unsafe.Sizeof(bool(true)))
	fmt.Printf("int32 size:%d\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int8 size:%d\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int64 size:%d\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("byte size:%d\n", unsafe.Sizeof(byte(0)))
	fmt.Printf("string size:%d\n", unsafe.Sizeof(string("0")))
	//align指的是对齐值,对齐准则尽量先放size小的
	//1.结构体成员变量，对齐值是默认对齐值（64位是8）和当前size的最小值，此时的偏移量必须是对齐值的整数倍
	//2.结构体本身的对齐值是默认对齐值和成员变量size中的最大值，最多为8
	//3.再次增加数据，只在对齐值之后进行，而不是实际使用字节之后
	//拿part1举例，1+4=5，必须是4的整数倍变成8，有三个空白的区域，下次内存存值的时候，从8开始
	//part1: 1,1+4~8,8+1~12,12+8~24,24+1~8*4=32;最终的内存占用32字节；对齐量是8
	//part2: 1+1=2,2+1=3,3+4~8,8+8=16；最终内存占用16字节，对齐量是8
	fmt.Printf("Part1 size:%d, align:%d \n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	fmt.Printf("Part2 size:%d, align:%d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
	fmt.Printf("Part3 size:%d, align:%d\n", unsafe.Sizeof(part3), unsafe.Alignof(part3))
	
	fmt.Println("=========================")
	var a interface{}
	fmt.Println(unsafe.Sizeof(a), unsafe.Alignof(a))
	a = 1
	fmt.Println(unsafe.Sizeof(a), unsafe.Alignof(a))
	a = "ok"
	fmt.Println(unsafe.Sizeof(a), unsafe.Alignof(a))

	b := test{data: 3, number: 2}
	fmt.Println(unsafe.Sizeof(b), unsafe.Alignof(b))

	c := test1{data:3, number: 2}
	fmt.Println(unsafe.Sizeof(c), unsafe.Alignof(c))

	st := stack{ data: 0,next: nil, size:1}
	st1 := stack{data: 1, next: &st, size:2}
	fmt.Println(unsafe.Sizeof(st), unsafe.Alignof(st))
	fmt.Println(unsafe.Sizeof(st1), unsafe.Alignof(st1))

	//指针占用的内存大小8个字节（或许是由于操作系统是64位的,设置的8字节)
	stt := &stack{data:0,next:nil,size:1}
	stt1 := &stack{data:1,next:stt, size:2}
	fmt.Println(unsafe.Sizeof(stt), unsafe.Alignof(stt))
	fmt.Println(unsafe.Sizeof(stt1), unsafe.Alignof(stt1))
}
