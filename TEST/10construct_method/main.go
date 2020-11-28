// package main

// import "fmt"

// func main() {
// 	p1 := newPerson("张三", "上海", 30)
// 	fmt.Printf("%#v,   %p\n", p1, &p1)
// 	p1.Dream()
// 	p1.LivePlace("北京")
// 	fmt.Printf("%#v\n", p1)
// 	p1.LiveChange("江苏")
// 	fmt.Printf("%#v\n", p1.city)
// }

// type Person struct {
// 	name, city string
// 	age        int8
// }

// func newPerson(name, city string, age int8) *Person {
// 	return &Person{name, city, age}
// }

// func (p Person) Dream() {
// 	fmt.Printf("%v has a great dream\n", p.name)
// }

// func (p *Person) LivePlace(city string) { //指针类型
// 	p.city = city
// }

// func (p Person) LiveChange(city string) { //值类型 拷贝
// 	p.city = city
// }

package main

import "fmt"

func main() {
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
	user := User{"小白", 19, 20}
	fmt.Printf("%#v\n", user)
}

//MyInt 将int定义为自定义MyInt类型
type MyInt int8

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

type User struct {
	//匿名字段，这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
	string
	int8
	int16
}
