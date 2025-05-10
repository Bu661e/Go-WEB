package controller

import "net/http"

func RegisterHomeRoutes() {
	http.HandleFunc("/home1", handleHome1)
	http.HandleFunc("/home2", handleHome2)
	http.HandleFunc("/home3", handleHome3)
	http.HandleFunc("/home4", handleHome4)

}

func handleHome1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Home 1!"))
}

func handleHome2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Home 2!"))
}

func handleHome3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Home 3!"))
}
func handleHome4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Home 4!"))
}
