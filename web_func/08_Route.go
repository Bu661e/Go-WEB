package web_func

import (
	"net/http"

	"github.com/Bu661e/Go-Web/controller"
)

func Test08_route() {
	controller.RegisterRoutes()
	http.ListenAndServe(":8080", nil)

}
