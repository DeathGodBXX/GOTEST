测试GO的原生自加函数和自己所编译的自加函数效率对比：

## `asm_go` file

```shell
$go tool compile -S math.go >> output # 只编译math.go文件
# $go build main.go -o math  会把math.s也包含进去，进行编译。
$go tool objdump -S math.o >> output
# 省去多余的import函数fmt等的编译信息和.文件信息
```
如果想编译运行，直接`go build`会自动把.s文件包含进去编译；再./file 执行文件

## `test` file

```shell
$go test -gcflags=-l -bench=. benchmem  #-gcflags=-l禁用内联

```
测试文件必须*_test.go结尾，想要测试的函数必须是这种形式TestXxxx()，BenchmarkXxxx()


## `src_c` file 

c语言反汇编命令：

```shell
$gcc -S stack_heap.c -o stack_heap.s -m32  #-m32表示32位的编译命令

```

