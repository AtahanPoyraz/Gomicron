package auth

import (
	"net/http"
)

//---[ HANDLER AUTH PROCESSING ]----------------------------------------------------------//

func HandleAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := GenerateToken()
		isAuthenticated := false

		cookie, err := r.Cookie("G0M1CR0N4UTHK3Y")
		if err == nil && cookie.Value == token.AuthToken {
			isAuthenticated = true
		}

		if !isAuthenticated {
			http.Error(w, "Authentication Failure", http.StatusUnauthorized)
			//http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}