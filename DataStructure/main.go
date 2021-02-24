package main

import (
	"DataStructure/ArrayList"
	"fmt"
)

func main() {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append(11)
	list.Append("x2")
	list.Append(12)
	list.Append("lsjfsj")
	list.Append([]byte{'s', 't', 'r', 'i', 'n', 'g'})

	for i := 0; i < 10; i++ {
		list.Insert(2, i)
	}
	fmt.Println(list.Get(2))
	list.Delete(2)
	list.Set(2, 100)
	list.Clear()
	fmt.Println(list)
}
