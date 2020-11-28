package main

import "fmt"

func main() {
	//匿名结构体
	var man struct {
		name string
		age  int
	}
	man.name = "张三"
	man.age = 17
	fmt.Printf("%#v\n", man)

	//指针类型结构体,new定义指针结构体，得到结构体指针
	fmt.Printf("\n\n=========指针类型结构体==========\n\n")
	var person = new(Person)
	person = &Person{name: "海王", city: "北京", age: 25}
	*person = Person{name: "海王", city: "北京", age: 25}
	fmt.Printf("type of person:%T,\nvalue of person:%#v\n", person, person)

	//定义结构体为了区分属性用';',初始化用','
	//指针类型支持person.name，在底层（*person）.name
	person.name = "女海王"
	person.city = "上海"
	person.age = 20
	fmt.Printf("type of person:%T,\nvalue of person:%#v\n", person, person)
	//直接定义初始化指针类型结构体
	person1 := Person{name: "男海王", city: "香港"}
	person2 := &Person{name: "王中王之海王大集结", age: 23}
	//分配空间并实例化，相当于new了一下，再实例化了
	fmt.Printf("type of person:%T,\nvalue of person:%#v\n", person1, person1)
	fmt.Printf("type of person:%T,\nvalue of person:%#v\n", person2, person2)
	fmt.Printf("Pointer person:%p\n", person)
	fmt.Printf("Pointer person:%p\n", person2)

	person3 := Person{"王中王之瘪王", "上头", 18}
	fmt.Printf("type of person:%T,\nvalue of person:%#v\n", person3, person3)

	fmt.Printf("\n\n==========内存布局=========\n\n")
	var person4 *Person
	Example(person4)

}
