package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	var wt Writer
	f := &FileX{path: "/home/deathgod/GO/src", filename: "IOREADER"}
	t := &Terminal{}
	wt = f
	wt.write()
	wt = t
	wt.write()

}

type Writer interface {
	write()
}

type Terminal struct {
}

type FileX struct {
	filename, path string
}

func (t *Terminal) write() {
	fmt.Printf("输出信息\n")
}

func (f *FileX) write() {
	filename := path.Join(f.path, f.filename)
	fileHandle, err := os.Open(filename)
	defer fileHandle.Close()
	if err != nil {
		fmt.Printf("Error occurs when Opening the file\n")
		return
	}
	_, errs := fileHandle.Write([]byte("Are you ok?"))
	if errs != nil {
		fmt.Printf("写入失败\n")
		fmt.Printf("%#v\n", errs)
	} else {
		fmt.Printf("写入成功\n")
	}

}
