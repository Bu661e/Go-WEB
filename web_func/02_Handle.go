package web_func

import (
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type AboutHandler struct{}

func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func Test02_handle() {
	// 1.创造一个服务器
	// (1) mux为 nil 时候，使用 DefaultServerMux
	// http.ListenAndServe(addr, mux) //DefaultServerMux

	// (2) Handler为 nil 时候，使用 DefaultServerMux
	s := &http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	// 2.注册Handler
	// (1) http.Handle 是把一个Handler对象注册到DefaultServerMux
	http.Handle("/hello", &HelloHandler{})
	http.Handle("/about", &AboutHandler{})

	// (2) http.HandleFunc 是把一个函数注册到DefaultServerMux
	// HandleFunc 是一个适配器，接受一个函数作为参数，返回一个 Handler 对象
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Home!"))
	})
	http.HandleFunc("/welcome", handleWelcome)

	// (3) 这就是(2)的底层实现
	// http.HandlerFunc(handleWelcome) -> 是 http.Handler 类型
	// 这里做了一个类型转换，把函数类型转换为 Handler 类型
	http.Handle("/go", http.HandlerFunc(handleWelcome))
	s.ListenAndServe()
}
