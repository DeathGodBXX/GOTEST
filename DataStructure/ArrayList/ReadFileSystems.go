package ArrayList

import (
	"io/ioutil"
)

//use Stack to store folders, files to store all content
//use another Stack to record current level,since it's possible to jump two level
//and no signal to show index skip
func GetAll(path string) ([]string, error) {
	st := NewStack()
	files := []string{}
	st.Push(path)
	for !st.isEmpty() {
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

	return files, nil
}
