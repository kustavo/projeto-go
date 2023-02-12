package middleware

import (
	nethttp "net/http"

	"github.com/kustavo/benchmark/go/application/interfaces"
	"github.com/kustavo/benchmark/go/infrastructure/http"
)

func AuthMiddleware(auth interfaces.Authentication, h nethttp.HandlerFunc) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		tokenString := http.RequestToken(r)
		_, err := auth.FetchAuth(r.Context(), tokenString)
		if err != nil {
			http.ResponseUnauthenticated(w, err, nil)
			return
		}
		h.ServeHTTP(w, r)
	}
}
