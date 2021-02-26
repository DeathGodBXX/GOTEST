package StackArray

import "errors"

//use []interface{} to construct Stack
type StackArray interface {
	Size() int                   //获取当前栈大小
	Clear()                      //清空
	Pop() (interface{}, error)   //出栈
	Push(data interface{}) error //入栈
	isFull() bool                //栈是否满
	isEmpty() bool               //栈是否空
}

type Stack struct {
	data        []interface{}
	maxsize     int
	currentsize int
}

//创建栈
func NewStack() (stack *Stack) {
	stack = new(Stack)
	stack.data = make([]interface{}, 0, 10)
	stack.maxsize = 10
	stack.currentsize = 0
	return
}

func (mystack *Stack) Size() int {
	return mystack.currentsize
}
func (mystack *Stack) Clear() {
	mystack.data = make([]interface{}, 0, 10)
	mystack.maxsize = 10
	mystack.currentsize = 0
}

func (mystack *Stack) Pop() (interface{}, error) {
	if mystack.isEmpty() {
		return nil, errors.New("栈为空，没有元素可以删除")
	}
	popvalue := mystack.data[mystack.currentsize-1]
	mystack.data = mystack.data[:mystack.currentsize-1]
	mystack.currentsize--
	return popvalue, nil
}

func (mystack *Stack) Push(data interface{}) error {
	if mystack.isFull() {
		return errors.New("栈是满的，不能再添加元素")
	}
	mystack.data = append(mystack.data, data)
	mystack.currentsize++
	return nil
}

func (mystack *Stack) isFull() bool {
	return mystack.currentsize >= mystack.maxsize
}

func (mystack *Stack) isEmpty() bool {
	return mystack.currentsize == 0
}
