package main

import (
	"fmt"
	"path"
)

func main1() {
	var tag int
	_, err := fmt.Scan(&tag)
	if err != nil {
		fmt.Printf("输入错误\n")
		return
	}
	fmt.Printf("%d\n", tag)
	for {
		fmt.Printf("Are you ok?\n")
		break
	}
}

func main() {
	pathx := "/home/deathgod"
	filename := "GO"
	ele := path.Join(pathx, filename) //文件名连接
	// ele := filepath.Join(pathx, filename)
	fmt.Printf("%#v\n", ele)
}
