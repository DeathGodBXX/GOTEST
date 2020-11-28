package main

import "testing"

var Result1 int

func BenchmarkAddNative1(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = int(i) + int(i)
	}
	Result1 = r
}

func BenchmarkAddAsm1(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = add(int(i), int(i))
	}
	Result1 = r
}

//goos: linux
// goarch: amd64
// pkg: COMPARE/asm_go
// BenchmarkAddNative-6   	1000000000	         0.258 ns/op	       0 B/op	       0 allocs/op
// BenchmarkAddAsm-6      	715940768	         1.60 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	COMPARE/asm_go	1.613s
//执行次数  单次迭代执行时间  单次迭代分配多少字节  单次迭代分配多少内存
//我们可以看到go原生的自加其实比使用汇编写的代码要快的多，
// 这是因为 Go 现在还不支持汇编函数内联，所以调用汇编函数执行自加会有一些函数调用的性能损耗，
// 所以自加汇编函数实现有更高的负载。
