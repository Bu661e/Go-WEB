package web_func

import (
	"fmt"
	"net/http"
)

func Test05_form() {

	http.HandleFunc("/form", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		// 1. 使用Form 返回map[string][]string类型，
		// 但是对于一个key，会返回多个value，其中包括form中的和url中的
		// 原因是form的格式是enctype="application/x-www-form-urlencoded"，相当于把form数据放到url中
		fmt.Fprintln(writer, request.Form)

		// 2. 使用FormValue 只包含form中的数据，不包含url中的数据
		fmt.Fprintln(writer, request.PostForm)

		// 3. 只用MultipartForm 能得到multipart kv对
		// 但是只能在Content-Type为multipart/form-data的时候使用
		request.ParseMultipartForm(1024)
		fmt.Fprintln(writer, request.MultipartForm)

	})
	http.ListenAndServe(":8080", nil)

}
