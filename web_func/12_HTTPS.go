package web_func

import (
	"net/http"
	"time"
)

// HTTPS优势
// 1. 请求多路复用：HTTPS可以同时处理多个请求，而HTTP只能处理一个请求。
// 2. Header压缩：HTTPS可以对HTTP请求和响应的Header进行压缩，减少传输数据的大小。
// 3. 默认安全
// 4. Server Push

func Test12_HTTPS() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 模拟处理请求的耗时操作
		time.Sleep(2 * time.Second)
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}
