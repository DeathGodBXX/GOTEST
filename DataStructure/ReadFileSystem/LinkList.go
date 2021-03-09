//LinkList Stack, LinkList Queue
// Get fs Tree by Stack, deep search
// same Level fs by Queue, wide search
package ReadFileSystem

import (
	"DataStructure/LinkList"
	"io/ioutil"
)

//use Stack to store folders
func FsByStack(path string) []string {
	st := LinkList.NewStack()
	st.Push(path)
	files := []string{}
	for !st.IsEmpty() {
		newpath, _ := st.Pop()
		path = newpath.(string)
		files = append(files, path) //储存当前文件夹
		fs, _ := ioutil.ReadDir(path)
		for _, fi := range fs { //键值循环
			if fi.IsDir() {
				st.Push(path + "/" + fi.Name()) //path记录着上一级文件夹
			} else {
				files = append(files, path+"/"+fi.Name())
			}
		}
	}
	return files
}

//use Queue to store all files include files and folders
func FsByQueue(path string) []string {
	qe := LinkList.NewQueue()
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
