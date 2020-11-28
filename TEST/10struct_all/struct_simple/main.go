package main

import (
	"fmt"
)

func main() {
	var a Newint
	var b Myint
	fmt.Printf("type of a:%T\n", a) //类型定义对于外部是只可见本类型，上一层类型不可见
	fmt.Printf("type of b:%T\n", b) //类型别名对于外部是可见的，包括上一层类型，编译的时候会被替换成上一层类型

	//结构体的定义与实例化
	//如果在不同的文件要一起编译

	var s int = 1 //1.定义的同时初始化
	var sss = 1   //2.类型推导
	ss := 1       //3.:相当于var,定义说明
	var ssss int  //4.先定义后初始化
	ssss = 1
	fmt.Println(s, ss, sss, ssss)

	//实例化
	var dog1 Dog = Dog{
		name:   "黄旭山",
		gender: "雄的",
		age:    3}
	fmt.Printf("\n%#v\n", dog1)

	var dog2 = Dog{
		name:   "黄旭山",
		gender: "雄的",
		age:    3}
	fmt.Printf("\n%#v\n", dog2)

	dog3 := Dog{name: "黄旭山", gender: "雄的", age: 3}
	fmt.Printf("\n%#v\n", dog3)

	var dog4 Dog
	dog4 = Dog{
		name: "游明卓",
		age:  1,
		// gender: "雄",
	}
	fmt.Printf("\n%#v\n", dog4)

	var dog Dog
	dog.age = 2
	dog.name = "习建航"
	dog.gender = "雄"
	fmt.Printf("\n%#v\n", dog)

	//=======匿名结构体===========
	//前面的都要指明结构体名,只能用.访问
	var doggy struct {
		name, gender string
		age          int
	}
	doggy.name = "小黄"
	doggy.gender = "雄的"
	// doggy = {name: "小白", gender: "兄的", age: 3}
	// var doggy1 struct{name:"小白" gender:"雄的" age:5}
	fmt.Printf("%#v\n", doggy)

	//指针类型的结构体
	PersonExample()
}
