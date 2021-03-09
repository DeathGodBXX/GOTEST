package LinkList

import "errors"

type QueueLinkList StackLinkList

type QueueIterator StackIterator

//只包含head,tail的队列
type Queue struct {
	head, tail *Node //结构体不能赋值为nil，但是成员变量可以赋值为nil
	size       int
}

func NewQueue() *Queue {
	qe := new(Queue) //必须要申请内存
	qe.head = nil
	qe.tail = nil
	qe.size = 0
	return qe
}

func (qe *Queue) Clear() {
	qe.head = nil
	qe.tail = nil
	qe.size = 0
}

//push to tail, pop from head
func (qe *Queue) Push(data interface{}) {
	newqe := new(Node)
	newqe.data = data
	newqe.next = nil
	//非首次插入，还需要修改当前尾部节点的指向和尾部的指向
	if qe.tail != nil {
		qe.tail.next = newqe
	}
	//首次插入，修改头部指向和尾部指向
	if qe.head == nil {
		qe.head = newqe
	}
	qe.tail = newqe
	qe.size++
}

func (qe *Queue) Pop() (interface{}, error) {
	if qe.IsEmpty() {
		return nil, errors.New("队列为空")
	}
	data := qe.head.data //必须要判断head是否为nil，这里判断nil的语句在IsEmpty()中
	qe.head = qe.head.next
	qe.size--
	return data, nil
}

func (qe *Queue) IsEmpty() bool {
	return qe.head == nil
}

func (qe *Queue) Size() int {
	return qe.size
}

func (qt *Queue) HasNext() bool {
	return !qt.IsEmpty()
}

func (qt *Queue) Next() (data interface{}) {
	if !qt.HasNext() {
		return nil
	}
	data = qt.head.data
	qt.head = qt.head.next
	return data
}

func (qt *Queue) Iterator() *Queue {
	newqe := new(Queue)
	// newqe = qt same reason with stack
	newqe.head = qt.head
	newqe.tail = qt.tail
	newqe.size = qt.size
	return newqe
}
