package main

import "fmt"

func main() {
	user1 := User{
		name: "小李",
		age:  18,
		Address: Address{
			province: "江苏",
			city:     "宜兴市",
		}, //最好添加逗号
	}
	fmt.Printf("%#v\n", user1)
	user1.Address.province = "香港"
	user1.city = "铜锣湾"
	fmt.Printf("%#v\n", user1)

	Example()

	dog := Dog{
		Feet: 4,
		Animal: &Animal{
			name: "小白",
		},
	}
	fmt.Printf("%#v\n", dog)
	dog.move()
	dog.scratch()
}
