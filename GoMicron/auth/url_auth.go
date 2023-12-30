package auth

import (
	"net/http"
)

//---[ URL AUTH ]---------------------------------------------------------------------------------------------------------------------------------//

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := GenerateToken()
		isAuthenticated := false

		cookie, err := r.Cookie("G0M1CR0N4UTHK3Y")
		if err == nil && cookie.Value == token.AuthToken {
			isAuthenticated = true
		}

		if !isAuthenticated {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

//---[ FILE PATH ]---------------------------------------------------------------------------------------------------------------------------------//

func ServeFile(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
}