package main

import (
	"DataStructure/ArrayList"
	"DataStructure/ReadFileSystem"
	"fmt"
)

func maintest() {
	path := "/home/deathgod/GO/src"
	files := ReadFileSystem.FsByStack(path)
	for _, f := range files {
		fmt.Println(f)
	}

	files1 := ReadFileSystem.FsByQueue(path)
	for _, f1 := range files1 {
		fmt.Println(f1)
	}

	files2 := ReadFileSystem.FsByStackArrayList(path)
	for _, f2 := range files2 {
		fmt.Println(f2)
	}

	fmt.Println(len(files))
	fmt.Println(len(files1))
	fmt.Println(len(files2))
}

func mainqueue() {
	qe := ArrayList.NewQueue()
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
}
