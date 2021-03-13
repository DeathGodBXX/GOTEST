package main

import (
	"DataStructure/ArrayList"
	"DataStructure/StackArray"
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

func main2() {
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

	qe := ArrayList.NewQueue()
	qe.Push(11)
	qe.Push(22)
	qe.Push(33)
	for qe.HasNext() {
		next, _ := qe.Next()
		fmt.Println(next)
	}
	fmt.Println(qe.Pop())
	qe.Clear()

	fmt.Println("++++++++++")

	qe.Push(11)
	qe.Push(22)
	qe.Push(33)

	for qe.HasNext() {
		fmt.Println(qe.Next())
	}
	fmt.Println(qe.Pop())

}

func main3() {
	st := StackArray.NewStack()
	st.Push(1)
	st.Push("ok")
	st.Push([]byte{'s', 't', 'r', 'i', 'n', 'g'})
	st.Push(2)
	st.Push(3)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
}

func main4() {
	st := ArrayList.NewStack()
	st.Push(1)
	st.Push("ok")
	st.Push([]byte{'s', 't', 'r', 'i', 'n', 'g'})
	st.Push(2)
	st.Push(3)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Size())
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	fmt.Println(st.Size())
	st.Push(66)
	st.Push(77)
	st.Push(88)
	fmt.Println("\n push later:")
	for st.HasNext() {
		val, _ := st.Next()
		fmt.Println(val)
	}
}
