// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func main() {

// }

// func main() {
// 	s1 := []Person{
// 		{name: "小白", age: 18},
// 	}
// 	var s2 []Person
// 	copy(s2, s1) //copy不能复制结构体,只能是特定的几种类型
// 	// s2 = s1
// 	fmt.Printf("s1:%#v,s2:%#v\n", s1, s2)

// 	t1 := make([]interface{}, 2, 4)
// 	// for i := 0; i < len(t1); i++ {
// 	// 	//for 循环的次数最多是len(t1)，否则会报错，也就是说遍历最多遍历到len(v-1)
// 	// 	t1[i] = s1
// 	// }
// 	// t2 := make([]interface{}, 1, 2)
// 	var t2 string
// 	fmt.Printf("请输入t2：\n")
// 	_, err := fmt.Scanf("%s\n", &t2)
// 	if err != nil {
// 		fmt.Printf("scanf出错!!!\n")
// 		return
// 	}
// 	// copy(t1, t2)
// 	fmt.Printf("t1:%#v\n,t2:%#v\n", t1, t2)
// }

// func main() {
// 	p1 := Person{name: "小白", age: 12}
// 	var p2 interface{}
// 	p2 = p1
// 	fmt.Printf("p1:%#v,p2:%#v\n", p1, p2)
// }

// type People interface {
// 	Speak(string) string
// }

// type Student struct{}

// func (stu *Student) Speak(think string) (talk string) {
// 	if think == "sb" {
// 		talk = "你是个大帅比"
// 	} else {
// 		talk = "您好"
// 	}
// 	return
// }

// func main() {
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }

// func main() {
// 	var x interface{}
// 	x = true
// 	fmt.Printf("%#v,%T\n", x, x)
// 	x = 8
// 	fmt.Printf("%#v,%T\n", x, x)
// 	x = "string"
// 	fmt.Printf("%#v,%T\n", x, x)

// }

// func main() {
// 	studentInfo := make(map[string]interface{}, 10)
// 	fmt.Printf("%d\n", len(studentInfo))
// 	studentInfo["name"] = "小白"
// 	studentInfo["city"] = "北京"
// 	studentInfo["age"] = 18
// 	fmt.Printf("%#v\n", studentInfo)
// 	fmt.Printf("%v\n", studentInfo)
// 	fmt.Printf("%d\n", len(studentInfo))

// }

// func main() {
// 	var t1 interface{}
// 	fmt.Printf("请输入字符串：\n")
// 	fmt.Scan(t1)
// 	fmt.Printf("%#v\n", t1)
// }

package main

import "fmt"

func main() {
	var x interface{}
	x = "hello,everyone!!!"
	v, ok := x.(string)
	if ok {
		fmt.Printf("%#v\n", v)
		// fmt.Printf("%v\n", x.(type))
	} else {
		fmt.Printf("x不是string类型的\n")
	}
	JustifyString(x)
}

func JustifyString(x interface{}) {
	switch v := x.(type) { //反射的使用
	case string:
		fmt.Printf("x is string type, %#v\n", v)
	case int:
		fmt.Printf("x is int type,%#v\n", v)
	case bool:
		fmt.Printf("x is bool type,%#v\n", v)
	}
}
