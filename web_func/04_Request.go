package web_func

import (
	"fmt"
	"net/http"
)

func request_basic() {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// 1.获取请求的方法
		fmt.Println("Method:", r.Method)
		// 2.获取请求的URL
		fmt.Println("URL:", r.URL)
		// 3.获取请求的协议版本
		fmt.Println("Proto:", r.Proto)
		// 4.获取请求的头部
		fmt.Println("Header:", r.Header)
		fmt.Println("Header:", r.Header["User-Agent"])
		fmt.Println("Header:", r.Header.Get("User-Agent"))
		// 5.获取请求的Body
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Println("Body:", string(body))
	})

}

// 对于GET请求，请求参数通常放在URL中，例如： /get?id=123&thread_id=456
func request_query() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		// 获取请求的参数
		query := r.URL.Query() // map[string][]string

		// 返回一个字符串数组，因为一个key可能对应多个value
		id := query["id"] // []string{"123", "456", "789"}
		fmt.Println("id:", id)

		// 返回第一个字符串
		threadID := query.Get("thread_id") // "123"
		fmt.Println("thread_id:", threadID)

	})
}

// TODO
func request_post_form() {

}

func Test04_request() {

	request_basic()
	request_query()
	request_post_form()
	http.ListenAndServe(":8080", nil)
}
