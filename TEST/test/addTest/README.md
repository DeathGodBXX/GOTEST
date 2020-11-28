编译命令：

```shell
#打印较长的完整编译信息（包括runtime的编译信息）和GO code对照，最后包含自己所写代码的汇编码信息
$go build main.go -o main
$go tool objdump -S main >> objdumpfile
#打印较短的编译信息（只输出自己所写的汇编代码信息，会生成编码文件，默认是main.o）
$go tool compile -S main.go >> objdumpfile 
$go tool objdump -S main.o >> objdumpfile
```

