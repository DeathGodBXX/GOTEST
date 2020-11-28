```shell
$ go run -gcflags '-m -l' main.go 
# -m 打印出优化策略
# -l 禁用内联
$ go tool compile -m -l main.go
# -m 打印出优化策略，相较于上一句命令，只打印出优化策略
```

