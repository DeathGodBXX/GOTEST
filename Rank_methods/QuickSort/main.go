package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 100000)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(1000000)
	}
	start := time.Now()
	QuickSort(a, 0, len(a)-1)
	fmt.Printf("collapsed of execution:%#v\n", float64(time.Since(start))/1e6)
	// fmt.Printf("%#v\n", a)
}

//l,r起着标志排序切片的区间段
func QuickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	key := arr[l]
	for { //交替执行，所以外部需要控制循环的前进
		if j <= i {
			break
		}
		for ; j > i; j-- { //先判断后执行,交换一次j,i的值，留出空位
			if arr[j] < key {
				arr[i] = arr[j]
				break
			}
		}
		for ; i < j; i++ { //先判断后执行，交换一次i,j的值，留出空位
			if arr[i] > key {
				arr[j] = arr[i]
				break
			}
		}
	}
	arr[i] = key
	QuickSort(arr, l, i-1) //不能在l，r的基础上修改左右偏移，l,r包含切片的区间段
	QuickSort(arr, i+1, r)
}

//快速排序
//分而治之的手段
//以第一个数据为标准，把小于此数的数据放在左边，大于此数的放在右边
//对于左右两边的数据同样的操作

//快速排序算法复杂度说明:
//分而治之
//可以这样思考，现将N元数组切割成logN片，类似于树结构
//而在第一层每个数据要进行遍历，不管从左到右，还是从右到左，最终都是遍历一整个数组，故而每一层都是N
//在往下排序都是遍历每个小区间段，加起来还是N,以此类推，每一层上都是N
//而一共有logN层，复杂度是O(N*logN)
