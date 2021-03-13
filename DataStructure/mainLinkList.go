package main

import (
	"DataStructure/LinkList"
	"DataStructure/ReadFileSystem"
	"fmt"
)

func main11() {
	st := LinkList.NewStack()
	st.Push(1)
	st.Push(2.1)
	st.Push("ok")

	for v, err := st.Pop(); err == nil; v, err = st.Pop() {
		fmt.Println(v)
	}
	println("st is empty? ", st.IsEmpty())
	for i := 0; i < 10; i++ {
		st.Push(i)
	}
	st.Clear()
	fmt.Println("st is empty? ", st.IsEmpty())

	st.Push(11)
	st.Push(22)
	st.Push(33)
	for it := st.Iterator(); it.HasNext(); {
		fmt.Println(it.Next(), st.Size())
	}
	fmt.Println(st.Pop())
}

func main22() {
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
	it := qe.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
	fmt.Println(qe.Pop())
}

func main33() {
	path := "/home/deathgod/GO/src"
	files := ReadFileSystem.FsByStack(path)
	for _, f := range files {
		fmt.Println(f)
	}

	path1 := "/home/deathgod/GO/src"
	files1 := ReadFileSystem.FsByQueue(path1)
	for _, f1 := range files1 {
		fmt.Println(f1)
	}
	fmt.Println(len(files))
	fmt.Println(len(files1))
}
