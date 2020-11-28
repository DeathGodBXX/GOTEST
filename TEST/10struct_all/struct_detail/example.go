package main

import "fmt"

type Person struct {
	name, city string
	age        int
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func Example(person *Person) *Person {
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	//即使把type放在Example里面也是逃逸的，有这种思路，只要觉得效率比较差的，可以写asm文件，设置成非逃逸的
	person = &Person{"海旺旺", "广东", 13}
	//person不逃逸，但是&Person逃逸，估计gc只要碰到取地址的都设置成逃逸的
	person.name = "海旺旺"
	person.city = "广东"
	person.age = 15
	return person
}

//执行命令：go tool compile -m -l *.go
