package web_func

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Bu661e/Go-Web/middleware"
)

func Test11_context() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			Name:    "Bu661e",
			Age:     18,
			Hobbies: []string{"coding", "reading", "running"},
		}

		time.Sleep(4 * time.Second)

		json.NewEncoder(w).Encode(c)
	})

	// 中间件在DefaultServerMux之前
	http.ListenAndServe(":8080", &middleware.TimeoutMiddleware{
		Next: new(middleware.AuthMiddleware),
	})
}
