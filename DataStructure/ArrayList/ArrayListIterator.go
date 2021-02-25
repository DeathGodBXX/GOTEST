package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool                              //是否有下一个
	Next() (interface{}, error)                 //获取当前索引下的结果
	Remove()                                    //删除当前元素
	GetIndex() int                              //获取当前索引
	Insert(index int, newval interface{}) error //插入一个元素
}

type Iterable interface {
	Iterator() Iterator
}

type ArrayListIterator struct {
	list         *ArrayList
	currentindex int
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentindex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentindex < it.list.theSize
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个元素")
	}
	value, err := it.list.Get(it.currentindex)
	it.currentindex++
	return value, err
}

func (it *ArrayListIterator) Remove() {
	it.list.Delete(it.currentindex)
}

func (it *ArrayListIterator) GetIndex() int {
	return it.currentindex
}

func (it *ArrayListIterator) Insert(index int, newval interface{}) error {
	return it.list.Insert(index, newval)
}
