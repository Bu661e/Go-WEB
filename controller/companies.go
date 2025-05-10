package controller

import (
	"net/http"
	"regexp"
)

func RegisterCompaniesRoutes() {
	http.HandleFunc("/companies", handleCompanies)
	http.HandleFunc("/companies/", handleCompany)
}
func handleCompany(w http.ResponseWriter, r *http.Request) {
	patten, _ := regexp.Compile(`/companies/(\d+)`)
	matches := patten.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		w.Write([]byte("Company " + matches[1]))
	} else {
		w.Write([]byte("Not Found"))
	}

}

func handleCompanies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Companies"))

}
