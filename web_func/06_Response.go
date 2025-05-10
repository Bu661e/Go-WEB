package web_func

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func example_write(writer http.ResponseWriter, request *http.Request) {
	str := `<html>
	<head>
		<title>Go Web</title>
	</head>
	<body>
		<h1>Hello World!</h1>
	</body>
	</html>`
	// writer.WriteHeader(200)
	writer.Write([]byte(str))
}

func example_writeHeader(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(501)
	fmt.Fprintf(writer, "Not Implemented")
}
func example_header(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Location", "/www.google.com")
	// 记住，在调用WriteHeader之前，修改Header，否则之后就不能修改了
	writer.WriteHeader(302)
}

// 定义 Post 结构体
type Post struct {
	User    string   `json:"user"`
	Threads []string `json:"threads"`
}

func example_json(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Bu661e",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	writer.Write(json)
}

func Test06_response() {

	s := &http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.HandleFunc("/write", example_write)
	http.HandleFunc("/writeHeader", example_writeHeader)
	http.HandleFunc("/hander", example_header)
	http.HandleFunc("/json", example_json)
	s.ListenAndServe()

}
