package SortAlgorithm

//单调递增奇偶排序(对比冒泡排序，此排序多重交换之后，最终肯定能将最大的数字往后移动和最小的数字往后移动)
func OESort(db []int) []int {
	if len(db) == 1 {
		return db
	}
	index := 0
	for {
		for i := index; i < len(db); i = i + 2 {
			if (i+1) < len(db) && db[i] > db[i+1] {
				db[i], db[i+1] = db[i+1], db[i]
			}
		}
		//利用tag标志跳出最终的死循环，同时跳出当前的for循环
		tag := true
		for j := 1; j < len(db); j++ {
			if db[j-1] > db[j] {
				tag = false
				break
			}
		}
		if tag {
			break
		}
		index = (index + 1) % 2
	}
	return db
}
