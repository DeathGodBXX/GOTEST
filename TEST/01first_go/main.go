package main

//在哪个文件夹下编译，可执行文件就放在那个文件下
//打开GO111MODULE="on",必须使用go mod init 包管理工具生成go.mod文件，否则package会显示错误，但运行正常
//优先级：go.mod,vendor,GOPATH/bin,由于1.11以后，go.mod成为默认的包管理工具
import (
	"01first_go/pack" //文件名称不重要，重要的是包名和方法名,同级目录下的pack文件夹，不需要replace
	"fmt"
	"test" //这个只是package代号,需要在go.mod中，利用replace指明路径“../02test_go”
)

func main() {
	fmt.Println("I'm handsome")
	pack.New()
	test.Old()
}
