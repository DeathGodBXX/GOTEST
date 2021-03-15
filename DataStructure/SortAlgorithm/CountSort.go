package SortAlgorithm

//单调递增且按次序的计数排序
//计数排序需要预先知道数的范围
func CountSort(db1 []string, db2 []int) []string {
	if len(db1) == 1 && len(db2) == 1 {
		return db1
	}
	//record max and min integer
	max := db2[0]
	min := db2[1]
	for _, v := range db2 {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	//辅助统计个数和相应的位置
	countArray := make([]int, max-min+1)
	for _, v := range db2 {
		countArray[v-min]++
	}
	for i := 1; i <= max-min; i++ {
		countArray[i] += countArray[i-1]
	}

	//记录指定索引对应的value
	result := make([]string, len(db1))
	for i := len(db2) - 1; i >= 0; i-- {
		result[countArray[db2[i]-min]-1] = db1[i]
		countArray[db2[i]-min]--
	}
	return result

}
