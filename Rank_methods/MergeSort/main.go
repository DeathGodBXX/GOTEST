package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, 100000)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1000000)
	}
	// fmt.Printf("%#v\n", array)
	start := time.Now()
	MergeSort(array, 0, len(array)-1)
	// fmt.Printf("=============================================\n")
	// fmt.Printf("%#v\n", array)
	fmt.Printf("collapsed of execution:%#v\n", float64(time.Since(start))/1e6)

}

func MergeSort(arr []int, l int, r int) {
	//无序时候的操作
	if l >= r {
		return
	}
	var s int
	s = int(math.Floor((float64(l+r) / 2)))
	MergeSort(arr, l, s)
	MergeSort(arr, s+1, r)
	//排序两个有序的slice
	AddTwo(arr, l, s+1, r)
}

func AddTwo(arr []int, left, share, right int) {
	result := make([]int, (right - left + 1))
	leftcopy, sharecopy := left, share
	for i := 0; i < len(result); i++ {
		// 保证一个数组耗尽，另一个数组直接全部复制到result中
		if leftcopy >= share {
			result = append(result[:i], arr[sharecopy:(right+1)]...) //已经执行了i++操作，左闭右开
			break
		} else if sharecopy > right {
			result = append(result[:i], arr[leftcopy:share]...)
			break
		}
		// 分别按次序把数组的元素放到result切片中
		if arr[leftcopy] < arr[sharecopy] {
			result[i] = arr[leftcopy]
			leftcopy++
		} else {
			result[i] = arr[sharecopy]
			sharecopy++
		}
	}
	copy(arr[left:(right+1)], result)
}

//注意事项：切片是左闭右开的
//传递的是两个有序切片的比较，[left:share],[share,right]
//一个切片耗尽之后，leftcopy必然大于等于share或者sharecopy大于right，这是由于share是达不到的，而rightkey达到

//如果要是无序的，切割切片，递归排序，这是树的后序遍历方式，往下进行切割，往上进行排序
//如果要是都有序，递归排序

//归并排序
//算法复杂度说明：
//切割N元数组，需要logN步，类似于树结构，底层两两比较，需要遍历N个元素,
//而每一层都需要两两比较，从而算法复杂度是O(N*logN)
