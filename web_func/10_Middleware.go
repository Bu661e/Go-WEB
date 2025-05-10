package web_func

import (
	"encoding/json"
	"net/http"

	"github.com/Bu661e/Go-Web/middleware"
)

func Test10_middleware() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			Name:    "Bu661e",
			Age:     18,
			Hobbies: []string{"coding", "reading", "running"},
		}
		json.NewEncoder(w).Encode(c)
	})

	// 中间件在DefaultServerMux之前
	http.ListenAndServe(":8080", new(middleware.AuthMiddleware))
}
