package ArrayList

import "errors"

//use ArrayList to construct Stack
type StackArrayList interface {
	Size() int                   //获取当前栈大小
	Clear()                      //清空
	Pop() (interface{}, error)   //出栈
	Push(data interface{}) error //入栈
	IsFull() bool                //栈是否满
	IsEmpty() bool               //栈是否空
}

type Stack struct {
	array   *ArrayList
	maxsize int
}

//创建栈
func NewStack() (stack *Stack) {
	stack = new(Stack)
	stack.array = NewArrayList()
	stack.maxsize = 1000
	return
}

func (mystack *Stack) Size() int {
	return mystack.array.Size()
}

func (mystack *Stack) Clear() {
	mystack.array = NewArrayList()
}

func (mystack *Stack) Pop() (interface{}, error) {
	if mystack.IsEmpty() {
		return nil, errors.New("栈为空，没有元素可以删除")
	}
	popvalue := mystack.array.dataStore[mystack.array.theSize-1]
	mystack.array.dataStore = mystack.array.dataStore[:mystack.array.theSize-1]
	mystack.array.theSize--
	return popvalue, nil
}

func (mystack *Stack) Push(data interface{}) error {
	if mystack.IsFull() {
		return errors.New("栈是满的，不能再添加元素")
	}
	mystack.array.dataStore = append(mystack.array.dataStore, data)
	mystack.array.theSize++
	return nil
}

func (mystack *Stack) IsFull() bool {
	return mystack.array.theSize >= mystack.maxsize
}

func (mystack *Stack) IsEmpty() bool {
	return mystack.array.theSize == 0
}
