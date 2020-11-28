package main

import "fmt"

type Dog struct {
	name   string
	age    int
	gender string
}
type Newint int

type Myint = int

type Person struct {
	name, city string
	age        int
}

func PersonExample() {
	fmt.Printf("\n\n==================\n\n")
	//new得到的是结构体地址，再实例化
	var person = new(Person)
	fmt.Printf("%T\n", person)
	fmt.Printf("%#v\n", person)
	person.name = "小白"
	person.city = "北京"
	fmt.Printf("%#v\n", person)

	fmt.Printf("\n\n=================\n\n")
	//取结构体的地址实例化
	p1 := &Person{"小黄", "上海", 23}
	p1.age=10
	fmt.Printf("%T\n", p1)
	fmt.Printf("%#v\n", p1)
	//可以看到两种结果是一样的，都可以通过p1.name取值，GO底层做的是（*p1）.name
}
