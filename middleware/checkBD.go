package middleware

import (
	"net/http"

	"github.com/godoquin/twittor/database"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
