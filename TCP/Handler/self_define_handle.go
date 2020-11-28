package main

import "net/http"

type helloHandler struct{}

func (mh *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("shopping, go shopping with me!!!"))
}

type aboutHandler struct{}

func (ab *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about, about you,do you know???"))

}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!!!!"))
}
func main() {
	//1.简单的连接
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { //响应数据
		w.Write([]byte("hello,everyone!Go is a beautiful programming language!!!"))
	})
	// http.HandleFunc("/welcome", welcome)
	http.Handle("/welcome", http.HandlerFunc(welcome))
	//上述两式等价
	//把welcome适配成Handler,HandlerFunc(f)强制类型转换，把welcome转换成HandlerFunc类型，
	//此类型有ServeHTTP方法（调用放f(w,r)）,装饰器

	//2.自定义Handler,处理网络请求
	// mh := &helloHandler{}
	// server := &http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: mh,
	// }
	// server.ListenAndServe()

	//3.将自定义的多个Handler注册到DefaultServeMux中（多路复用器） 适用于逻辑复杂的handler
	mh := &helloHandler{}
	http.Handle("/shopping", mh)
	ab := &aboutHandler{}
	http.Handle("/about", ab)
	http.ListenAndServe("localhost:8081", nil)

	//HandleFunc第二个参数是Handler函数，适配成handler，更自由；
	//Handle第二个参数是handler，可以组合handler的serveHTTP方法，新定义接口，补充现有功能

}
