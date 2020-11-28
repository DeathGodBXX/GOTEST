package main

import "fmt"

func main() {
	s := make([]Student, 0, 100)
	// s := newStudent(12, 18, 90, "小白")
	// fmt.Printf("%#v\n", s)
	for {
		fmt.Printf("请选择选项：\n1.添加学生\n2.修改学生列表\n3.展示所有学生信息\n4.删除指定学生\n5.退出\n")
		var n int
		fmt.Scanf("%d\n", n)
		switch n {
		case 1:
			s = addStudent(s)
		case 2:
			// changeStudent()
		case 3:
			// showStudent()
		case 4:
			// deleteStudent()
		case 5:
			break
		default:
			fmt.Printf("输入错误，重新操作\n")
		}

	}
}

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

type StudentList struct {
	Size     int
	Students []Student
}

// func newStudent(Id, Age, Score int, Name string) *Student {
// 	return &Student{
// 		Id:    Id,
// 		Name:  Name,
// 		Age:   Age,
// 		Score: Score,
// 	}
// }

//添加学生信息
func addStudent(s []Student) []Student {
	for {
		fmt.Printf("请按照id,name,score,age的顺序输入信息：\n")
		var id, score, age int
		var name string
		_, err := fmt.Scanf("%d %s %d %d", id, name, score, age)
		if err != nil {
			fmt.Printf("类型输入错误，请重新输入\n")
			continue
		}
		return append(s, Student{Id: id, Name: name, Score: score, Age: age})
	}
}

//展示学生列表
// func showStudent(s []Student) {
// 	for _, v := range s {
// 		fmt.Printf("%#v\n", v)
// 	}
// }
func showStudent(s []Student) {
	// fmt.Printf("%d\n", len(*s))
	// for i := 0; i < len(s); i++ {

	// }
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%v\n",s[i])
	// }
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}

//编辑学生信息
// func (s *Student) changeStudent() {
// 	fmt.Printf("请输入想要改变的选项:\n1.name\n2.age\n3.score\n4.id\n")
// 	var n int
// 	fmt.Scanf("%d\n", n)
// 	switch n {
// 	case 1:
// 		fmt.Printf("请输入想要修改的姓名:")
// 		var name string
// 		fmt.Scanf("%s\n", name)
// 		//查询姓名
// 		for _,v = *s{
// 			if name == v.Name
// 		}
// 	case 2:
// 	case 3:
// 	case 4:
// 	default:
// 		fmt.Printf("输入错误，请重新输入：")
// 	}

// }

func deleteStudent(s []Student) { //追加根据id删除信息
	fmt.Printf("请输入想要删除的学生姓名:\n")
	var name string
	fmt.Scanf("%s\n", name)
	// fmt.Printf("%d\n", len(s))
	tag := false
	for i, v := range s {
		if name == v.Name {
			s = append(s[:i], s[(i+1):]...)
			tag = true
		}
	}
	if tag == false {
		fmt.Printf("不存在这个人的信息\n")
	}

}
