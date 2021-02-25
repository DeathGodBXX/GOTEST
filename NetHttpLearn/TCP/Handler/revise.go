package main

import "net/http"

func main() {
	//1.simple use
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello,world!!"))
	// })
	// http.ListenAndServe(":8080", nil)

	http.HandleFunc("/home", home) //把func注册成handler

	http.Handle("/rest", http.HandlerFunc(rest)) //把func注册成handler

	mh := &myHandler{}
	http.Handle("/myhandler", mh)

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	server.ListenAndServe()

}

func home(w http.ResponseWriter, r *http.Request) { //自定义
	w.Write([]byte("home page"))
}

func rest(w http.ResponseWriter, r *http.Request) { //自定义
	w.Write([]byte("rest page"))
}

type myHandler struct{} //组合Handler,构成新的Handler
func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myhandler constructor"))
}
