package main

import (
	"DataStructure/LinkList"
	"fmt"
)

func mainst() {
	st := LinkList.NewStack()
	st.Push(1)
	st.Push(2.1)
	st.Push("ok")

	for v, err := st.Pop(); err == nil; v, err = st.Pop() {
		fmt.Println(v)
	}
	println("st is empty? ", st.Isempty())
	for i := 0; i < 10; i++ {
		st.Push(i)
	}
	st.Clear()
	fmt.Println("st is empty? ", st.Isempty())

	st.Push(11)
	st.Push(22)
	st.Push(33)
	for it := st.Iterator(); it.HasNext(); {
		fmt.Println(it.Next(), st.Size())
	}
	fmt.Println(st.Pop())
}

func main() {
	qe := LinkList.NewQueue()
	qe.Push(1)
	fmt.Println(qe.Size())
	qe.Push("ok")
	fmt.Println(qe.Size())
	qe.Push(2.2)
	fmt.Println(qe.Size())

	fmt.Println(qe.Pop())
	fmt.Println(qe.Size())
	qe.Clear()
	fmt.Println(qe.Size())
	fmt.Println(qe.IsEmpty())

	qe.Push(11)
	qe.Push(22)
	qe.Push(33)
	for it := qe.Iterator(); it.HasNext(); {
		fmt.Println(it.Next(), qe.Size())
	}
	fmt.Println(qe.Pop())
}

// func main() {
// 	path := "/home/deathgod/Go/src"
// 	//利用迭代器模式设置栈迭代器，队列迭代器
// 	// for v := range LinkList.GetFs(path) {

// 	// }
// }
