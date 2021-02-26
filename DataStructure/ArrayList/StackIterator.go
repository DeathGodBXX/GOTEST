package ArrayList

import "errors"

// use Stack to construct StackIterator
type IteratorStack interface {
	HasNext() bool
	Next() (interface{}, error)
}

func (mystack *Stack) HasNext() bool {
	return mystack.array.theSize > 0
}

func (mystack *Stack) Next() (interface{}, error) {
	if !mystack.HasNext() {
		return nil, errors.New("没有下一个元素")
	}
	next, _ := mystack.array.Get(mystack.array.theSize - 1)
	mystack.array.theSize--
	return next, nil
}
