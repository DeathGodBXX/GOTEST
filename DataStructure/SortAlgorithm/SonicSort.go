package SortAlgorithm

//单调递增快速排序(递归方式，自己调用自己)，注意写成范型，需要类型断言，没有类是于cpp模板编程的单态化
//负责处理逻辑，调用细节(切片是引用类型)
func SonicSort(db []int) []int {
	if len(db) == 1 {
		return db
	}
	// index := partition(db)
	// //特殊情况处理选取的pivot位于首尾，容错处理
	// if index > 1 {
	// 	SonicSort(db[:index])
	// }
	// if index < len(db)-2 {
	// 	SonicSort(db[(index + 1):])
	// }
	index := partitionDP(db)
	if index > 1 {
		SonicSort(db[:index])
	}
	if index < len(db)-2 {
		SonicSort(db[(index + 1):])
	}
	return db
}

//负责每次有效搜索，pivot对应的索引（挖坑法,双指针单移动）
func partition(db []int) int {
	pivot := db[0]
	left, right := 0, len(db)-1
	tag := true
	for right != left {
		//利用tag标志是在右指针移动还是左指针移动，穷尽左指针和右指针
		//每次只有一个指针在移动
		if tag {
			if db[right] > pivot {
				right--
				continue
			}
			tag = false
			db[left] = db[right]
			left++
		} else {
			if db[left] < pivot {
				left++
				continue
			}
			tag = true
			db[right] = db[left]
			right--
		}
	}
	db[left] = pivot
	return left
}

//指针交换法
func partitionDP(db []int) int {
	pivot := db[0]
	left, right := 1, len(db)-1 //len(db)>=2
	for right > left {
		if db[right] > pivot {
			right--
			continue
		}
		if db[left] < pivot {
			left++
			continue
		}
		//对于重复出现的pivot,不需要特殊处理，这次只是定位db[0]的位置，下次分治法最多也是以同一值做分治。
		db[right], db[left] = db[left], db[right]
		left++
		right--
	}
	//防止dp[0]之后的元素比dp[right]小，误交换
	//right==left,判断当前节点与pivot大小，如果较大(归属于right区域)，则将db[0]与db[right-1]交换,返回right-1；
	//如果小于等于，则把db[0]和db[right]交换(归属于left区域),返回right；
	//right<left，此时，right指针跑道left区域最后一个元素(归属于left区域),执行db[right]和db[0]的交换,返回right

	if db[0] < db[right] {
		db[right-1], db[0] = db[0], db[right-1]
		return right - 1
	}
	db[right], db[0] = db[0], db[right]
	return right
}
