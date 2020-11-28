package main

import "fmt"

func main() {
	studentA := new(StudentArray)
	// s := make([]Student, 0, 100)
	studentA.Students = make([]Student, 0, 100)
BREAKELEM:
	for {
		fmt.Printf("\n\n=======================\nManipulation:\n1.Modify\n2.Add\n3.Show All\n4.Delete\n5.Exit\n==================\n\nChoice: ")
		var tag int
		_, err := fmt.Scan(&tag) //interface类型如何从标准输入读取信息
		if err != nil {
			fmt.Printf("Enter wrong tag,please choose again!!!!\n")
			continue
		}
		switch tag {
		case 1:
			studentA.ModifyStudent()
		case 2:
			studentA.AddStudent()
		case 3:
			studentA.ShowStudent()
		case 4:
			studentA.DeleteStudent()
		default:
			break BREAKELEM
		}
	}
}

type Student struct {
	Id, Age, Score int
	Name           string
}

type StudentArray struct {
	Size     int
	Students []Student
}

//Modify info
func (s *StudentArray) ModifyStudent() {
	fmt.Printf("Enter student's name:")
	var name string
	fmt.Scanln(&name)
	var i int
	for i = 0; i < s.Size; i++ {
		if name == s.Students[i].Name {
			fmt.Printf("Enter new infomation:id,age,score\n")
			fmt.Scanln(&s.Students[i].Id, &s.Students[i].Age, &s.Students[i].Score)
			break
		}
	}
	if i == s.Size {
		fmt.Printf("No this character!!!!!\n")
	}
}

//add info
func (s *StudentArray) AddStudent() {
	fmt.Printf("Enter full infomation:name,id,age,score\n")
	var std Student
	fmt.Scanln(&std.Name, &std.Id, &std.Age, &std.Score)
	var flag = false
	for i := 0; i < s.Size; i++ {
		if std.Name == s.Students[i].Name {
			fmt.Printf("The name is in the StudentArray,you maybe input wrong name,please check again!!!!!\n Or you maybe Modify the existing Information,choose step 1!!!!!\n")
			flag = true
			break
		}
	}
	if !flag {
		s.Students = append(s.Students, std)
		s.Size++
	}
}

//show all info
func (s *StudentArray) ShowStudent() {
	for i := 0; i < s.Size; i++ {
		fmt.Printf("%#v\n", s.Students[i])
	}
}

//delete someone info
func (s *StudentArray) DeleteStudent() {
	fmt.Printf("Enter student's name:\n")
	var name string
	fmt.Scanln(&name)
	tag := true
	for i := 0; i < s.Size; i++ {
		if name == s.Students[i].Name {
			s.Students = append(s.Students[:i], s.Students[(i+1):]...)
			s.Size--
			tag = false
		}
	}
	if tag {
		fmt.Printf("No this character!!!!\n")
	}
}
