package SortAlgorithm

//单调递增堆排序
func HeapSort(db []int) []int {
	if len(db) == 1 {
		return db
	}
	for i := 0; i < len(db); i++ {
		heapSortMax(db[:len(db)-i])
	}
	return db
}

//将最大的元素登顶（上浮成最大堆）
func heapSortMax(db []int) []int {
	depth := len(db)/2 - 1
	for i := depth; i >= 0; i-- {
		//处理parent,2*parent+1,2*parent+2,topmax用于记录最大值对应的索引
		topmax := i
		if db[2*i+1] > db[topmax] {
			topmax = 2*i + 1
		}
		if (2*i+2) < len(db)-1 && db[2*i+2] > db[topmax] {
			topmax = 2*i + 2
		}
		if topmax != i {
			db[topmax], db[i] = db[i], db[topmax]
		}
	}
	db[0], db[len(db)-1] = db[len(db)-1], db[0]
	return db
}
