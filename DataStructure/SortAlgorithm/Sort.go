package SortAlgorithm

//go 有编译阶段，会把dbp单态化成[]int,[]string;还未支持[]interface{}作为参数传递
//单调递增选择排序
func SelectSort(dbp interface{}) interface{} {
	db := dbp.([]int)
	for i := 0; i < len(db)-1; i++ {
		//利用min缓存最小值
		min := i
		for j := i + 1; j < len(db); j++ {
			if db[j] < db[min] {
				min = j
			}
		}
		if min != i {
			db[min], db[i] = db[i], db[min]
		}
	}
	return db
}

//单调递增插入排序
func InsertSort(dbp interface{}) interface{} {
	db := dbp.([]int)
	for i := 1; i < len(db); i++ {
		ins := i
		//找出插入的位置
		for j := 0; j < i; j++ {
			if db[j] > db[i] {
				ins = j
				break
			}
		}
		if ins != i {
			// record := append(db[i:(i+1)], db[ins:i]...) //产生扩容或者覆盖了之前的数值
			// fmt.Println("record ", record)
			// record1 := append(db[:ins], record...)
			// fmt.Println("record1 ", record1)
			// db = append(record1, db[(i+1):]...)

			// for k := ins; k < i; k++ { //降低空间复杂度
			// 	db[k], db[i] = db[i], db[k]
			// }
			//相较于上一种交换的方法，减少读写内存的次数，同时降低空间复杂度
			record := db[i]
			for k := i - 1; k >= ins; k-- {
				db[k+1] = db[k]
			}
			db[ins] = record
		}
	}
	return db
}

//单调递增冒泡排序
func BubbleSort(dbp interface{}) interface{} {
	db := dbp.([]int)
	for i := len(db); i > 1; i-- {
		isneed := false
		for j := 0; j < i-1; j++ {
			if db[j] > db[j+1] {
				db[j], db[j+1] = db[j+1], db[j]
				isneed = true
			}
		}
		if !isneed {
			break
		}
	}
	return db
}
