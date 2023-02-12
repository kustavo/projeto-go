package middleware

import "net/http"

func CorsMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")
		allowed := []string{
			origin,
		}

		for _, val := range allowed {
			if val == origin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
		}

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
