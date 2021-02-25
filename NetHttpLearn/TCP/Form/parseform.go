package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http.HandleFunc("/getbody", func(w http.ResponseWriter, r *http.Request) {
	// 	length := r.ContentLength
	// 	body := make([]byte, length)
	// 	r.Body.Read(body)
	// 	fmt.Fprintln(w, string(body))
	// })

	// http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
	// 	//PostForm只支持application,MultipartForm支持form-data,而且能接收上传的文件，
	// 	//都是只包含表单数据，但是Form能包含url中数据
	// 	r.ParseMultipartForm(1024)
	// 	fmt.Fprintln(w, r.MultipartForm)
	// 	// fmt.Fprintln(w, r.FormValue("first_name"))
	// })

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// r.ParseMultipartForm(1 >> 10)
		// file := r.MultipartForm.File["upload"][0]
		// file1, err := file.Open()

		file1, _, err := r.FormFile("upload") //返回并打开文件
		defer file1.Close()
		if err == nil {
			data, err := ioutil.ReadAll(file1)
			if err == nil {
				fmt.Fprintln(w, string(data))
			}

		}
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		post := &Post{
			User: "xiaoming",
			Age:  20,
		}
		json, _ := json.Marshal(post)
		w.Write(json)
	})

	http.Handle("/", http.FileServer(http.Dir("Handler/wwwroot")))
	http.ListenAndServe(":8080", nil)
}

type Post struct { //注意结构体的变量首字母要大写，否则外部无法访问user,age，json序列化出错
	User string
	Age  int64
}
