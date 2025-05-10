package web_func

import (
	"net/http"
)

func Test03_insideHandler() {
	// 1.手写FileServer
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "wwwroot"+r.URL.Path)
	// })
	// http.ListenAndServe(":8080", nil)

	// 2.使用http内置的FileServer
	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))

}
