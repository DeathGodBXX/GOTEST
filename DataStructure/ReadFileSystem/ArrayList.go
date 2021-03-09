//use StackArrayList to store folders
//use QueueArrayList to store all files
package ReadFileSystem

import (
	"DataStructure/ArrayList"
	"io/ioutil"
)

func FsByStackArrayList(path string) []string {
	st := ArrayList.NewStack()
	files := []string{}
	st.Push(path)
	for !st.IsEmpty() {
		newpath, _ := st.Pop()
		path = newpath.(string)
		files = append(files, path) //存储当前处理的文件路径
		fs, _ := ioutil.ReadDir(path)
		for _, fi := range fs {
			if fi.IsDir() {
				st.Push(path + "/" + fi.Name())
				//不可以存储path = path + "/" + fi.Name(),否则当前文件夹下文件会连成一串
			} else {
				files = append(files, path+"/"+fi.Name())
			}
		}
	}
	return files
}

func FsByQueueArrayList(path string) []string {
	qe := ArrayList.NewQueue()
	qe.Push(path)
	files := []string{}
	for !qe.IsEmpty() {
		newpath, _ := qe.Pop()
		path = newpath.(string)
		files = append(files, path)
		fs, err := ioutil.ReadDir(path)
		if err == nil {
			for _, fi := range fs {
				qe.Push(path + "/" + fi.Name())
			}
		}
	}
	return files
}
