package ArrayList

import "errors"

type QueueIterator IteratorStack

func (qe *Queue) HasNext() bool {
	return qe.array.theSize > 0
}

func (qe *Queue) Next() (interface{}, error) {
	if !qe.HasNext() {
		return nil, errors.New("没有下一个元素")
	}
	next, _ := qe.array.Get(0)
	qe.array.dataStore = qe.array.dataStore[1:]
	qe.array.theSize--
	return next, nil
}

// func (qe *Queue) Iterator() *Queue {
// 	newqe := new(Queue)
// 	newqe.array = qe.array
// 	newqe.maxsize = qe.maxsize
// 	return newqe
// }
