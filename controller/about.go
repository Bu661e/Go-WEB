package controller

import (
	"net/http"
)

func RegisterAboutRoutes() {
	http.HandleFunc("/about1", handleAbout1)
	http.HandleFunc("/about2", handleAbout2)
	http.HandleFunc("/about3", handleAbout3)
	http.HandleFunc("/about4", handleAbout4)

}

func handleAbout1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello About 1!"))
}

func handleAbout2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello About 2!"))
}

func handleAbout3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello About 3!"))
}
func handleAbout4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello About 4!"))
}
