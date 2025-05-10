package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
	next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if am.next == nil {
		am.next = http.DefaultServeMux
	}

	auth := r.Header.Get("Authorization")
	if auth != "" {
		am.next.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
	}

}
