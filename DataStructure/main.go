package main

import (
	"DataStructure/ArrayList"
	"fmt"
)

func main1() {
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
	for it := list.Iterator(); it.HasNext(); {
		currentValue, _ := it.Next()
		fmt.Println(currentValue)
	}
}

func main() {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append(11)
	list.Append("x2")
	list.Append(12)
	list.Append("lsjfsj")
	list.Append([]byte{'s', 't', 'r', 'i', 'n', 'g'})

	fmt.Println(list)

	for it := list.Iterator(); it.HasNext(); {
		if it.GetIndex() == 2 {
			it.Remove() //append连接之后，仍然是连续存储的数组
		}
		fmt.Println(it.Next()) //Next()输出当前元素，currentindex++
	}

	list.Insert(2, "x2")
	list.Iterator().Insert(2, "y2")
	for it := list.Iterator(); it.HasNext(); {
		value, _ := it.Next()
		fmt.Println(value)
	}

	fmt.Println(list)

}
