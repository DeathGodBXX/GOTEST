package SortAlgorithm

import (
	"math"
	"sort"
)

func BucketSort(db []float64) []float64 {
	if len(db) == 1 {
		return db
	}
	//获取范围
	max := db[0]
	min := db[1]
	for i := 0; i < len(db); i++ {
		if db[i] > max {
			max = db[i]
		}
		if db[i] < min {
			min = db[i]
		}
	}
	//桶添加数据
	bucket := make([][]float64, len(db))
	for i := 0; i < len(db); i++ {
		if db[i] == max {
			bucket[len(db)-1] = append(bucket[len(db)-1], db[i])
			continue
		}
		index := int(math.Floor((db[i] - min) * float64(len(db)-1) / (max - min)))
		bucket[index] = append(bucket[index], db[i])
	}

	//归类到db中
	j := 0
	for i := 0; i < len(db); i++ {
		sort.Float64s(bucket[i])
		for _, v := range bucket[i] {
			db[j] = v
			j++
		}
	}
	return db
}
