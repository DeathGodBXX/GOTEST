package main

import "fmt"

type Address struct {
	province, city string
}

type User struct {
	name    string
	age     int8
	Address // Address 你名字段
}

//结构体嵌套
//Address 地址结构体
type Address1 struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User1 struct {
	Name   string
	Gender string
	Address1
	Email
}

func Example() {
	var user3 User1
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address1.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"    //指定Email结构体中的CreateTime
	fmt.Printf("%#v\n", user3)

}

//结构体方法继承
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%v can move\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //匿名结构体
}

func (b *Dog) scratch() {
	fmt.Printf("%v has %v feet\n", b.Animal.name, b.Feet)
}
