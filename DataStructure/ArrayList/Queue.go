package ArrayList

import "errors"

type QueueArrayList StackArrayList

type Queue Stack

func NewQueue() *Queue {
	qe := new(Queue)
	qe.array = NewArrayList()
	qe.maxsize = 1000
	return qe
}

func (qe *Queue) Size() int {
	return qe.array.theSize
}

func (qe *Queue) Clear() {
	qe.array = NewArrayList()
	qe.maxsize = 1000
}

func (qe *Queue) Pop() (interface{}, error) {
	if qe.IsEmpty() {
		return nil, errors.New("队列是空的，不能出队")
	}
	data := qe.array.dataStore[qe.array.theSize-1]
	qe.array.dataStore = qe.array.dataStore[:qe.array.theSize-1]
	qe.array.theSize--
	return data, nil
}

func (qe *Queue) Push(data interface{}) error {
	if qe.IsFull() {
		return errors.New("队列已满，不能入队")
	}
	qe.array.dataStore = append(qe.array.dataStore, data)
	qe.array.theSize++
	return nil
}

func (qe *Queue) IsFull() bool {
	return qe.array.theSize == qe.maxsize
}

func (qe *Queue) IsEmpty() bool {
	return qe.array.theSize == 0
}
