package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", welcome)

	//五种默认的内置handler
	// http.Handle("/", http.NotFoundHandler())
	http.Handle("/redirect", http.RedirectHandler("/welcome",
		http.StatusMovedPermanently))
	http.Handle("/heihei/welcome", http.StripPrefix("/heihei", http.HandlerFunc(welcome)))
	http.Handle("/home", http.TimeoutHandler(http.HandlerFunc(home),
		1<<30, "超时警告"))
	// http.HandleFunc("/fs/", func(w http.ResponseWriter, r *http.Request) {
	// 	//只有/fs，如果输入/fs/index.html,只会带着/fs/进入此Handler,截取文件名
	// 	fmt.Printf("%v\n", r.URL.Path[3:])
	// 	http.ServeFile(w, r, "wwwroot"+r.URL.Path[3:])
	// })
	http.Handle("/", http.FileServer(http.Dir("wwwroot")))
	//FileServer的前缀必须是根目录

	http.ListenAndServe(":8080", nil) //DefaultServeMux也是某种Handler

	// http.ListenAndServe(":8081", http.FileServer(http.Dir("wwwroot"))) //注册新的handler

}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome,everyone!!!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page success!!"))
}
