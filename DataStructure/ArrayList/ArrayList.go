package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int                                  //获取数组大小
	Get(index int) (interface{}, error)         //获取指定索引的值
	Set(index int, newval interface{}) error    //修改数据
	Insert(index int, newval interface{}) error //插入数据
	Append(newval interface{}) error            //追加数据
	Delete(index int)                           //删除指定位置的数据
	Clear()                                     //清空数据
	String() string                             //格式化输出结构体指定成员变量
}

type ArrayList struct {
	dataStore []interface{}
	theSize   int
}

//创建一个ArrayList
func NewArrayList() (list *ArrayList) {
	list = new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
	return
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Append(newval interface{}) error {
	list.dataStore = append(list.dataStore, newval)
	list.theSize++
	return nil
}

func (list *ArrayList) Size() int {
	return list.theSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界，get不到下一个")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("Set索引越界")
	}
	list.dataStore[index] = newval
	return nil
}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index > list.theSize {
		return errors.New("Insert索引越界")
	}

	if list.theSize == cap(list.dataStore) {
		newdataStore := make([]interface{}, 2*list.theSize, 2*list.theSize) //size是实际申请的内存
		copy(newdataStore, list.dataStore)
		list.dataStore = newdataStore
	}
	list.dataStore = list.dataStore[:list.theSize+1]
	//内存必须移动一位，手动设置，如果省去这一行内存size==list.theSize,一直导致无法移动插入新数据

	copy(list.dataStore[index+1:], list.dataStore[index:list.theSize])
	list.dataStore[index] = newval
	list.theSize++
	return nil
}

func (list *ArrayList) Delete(index int) {
	if index < 0 || index >= list.theSize {
		return
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.theSize--
	return
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
}
