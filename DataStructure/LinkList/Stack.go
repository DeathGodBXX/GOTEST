package LinkList

import "errors"

//LinkList realize stack
type StackLinkList interface {
	Push(data interface{})
	Pop() (interface{}, error)
	IsEmpty() bool
	Clear()
	Size() int
}

type Node struct {
	data interface{} //大小是16字节，对齐量是8字节
	next *Node
}

type Stack struct {
	top  *Node
	size int
}

//Iterator pattern
type StackIterator interface {
	HasNext() bool
	Next() interface{}
}

func NewStack() *Stack {
	st := new(Stack) //必须要先new一段内存
	st.top = nil
	st.size = 0
	return st
}

func (st *Stack) Push(data interface{}) {
	newstack := new(Node)
	newstack.data = data
	newstack.next = st.top
	st.top = newstack
	st.size++
}

func (st *Stack) Pop() (interface{}, error) {
	if st.Isempty() {
		return nil, errors.New("栈为空")
	}
	data := st.top.data
	st.top = st.top.next
	st.size--
	return data, nil
}

func (st *Stack) Isempty() bool {
	return st.top == nil
}

func (st *Stack) Clear() {
	st.top = nil
	st.size = 0
}

func (st *Stack) Size() int {
	return st.size
}

func (it *Stack) HasNext() bool {
	return !it.Isempty()
}

func (it *Stack) Next() (data interface{}) {
	if !it.HasNext() {
		return nil
	}
	data = it.top.data
	it.top = it.top.next
	it.size--
	return data
}

func (st *Stack) Iterator() *Stack {
	it := new(Stack)
	// it = st  it和st指向共用同一个指向Stack指针
	// 下面的赋值，it拷贝了st指向的Stack top和size,st指向的Stack不会改变，相当于深拷贝了1个Stack;
	//后续st可以继续使用（指向未被释放）
	it.top = st.top
	it.size = st.size
	return it
}
