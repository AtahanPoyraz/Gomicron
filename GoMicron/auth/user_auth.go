package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AtahanPoyraz/cmd"
)

//---[ USER AUTH ]-----------------------------------------------------------------------------------------------------------------------------------//

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) AuthUser(w http.ResponseWriter, r *http.Request) {
	userEmail := "atahanpoyraz@gmail.com"
	userPassword := "changeme"
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Error decoding JSON request body", http.StatusBadRequest)
		return
	}
	
	email := credentials.Email
	password := credentials.Password

	if email != userEmail || password != userPassword {
		http.Error(w, "Incorrect email or password. Please try again.", http.StatusBadRequest)
		return
	}
	
	if email == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}
	
	if email == userEmail && password == userPassword {
		token := GenerateToken()
		response := map[string]interface{}{
			"message":       "Authentication successful",
			"authenticated": true,
			"token":         token.AuthToken,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		s.l.Printf("%s[AUTH]%s : Authentication successful for user: %s", cmd.BYELLOW_BLACK, cmd.TRESET, email)
		return
	}

	response := map[string]interface{}{
		"message":       "Authentication unsuccessful",
		"authenticated": false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	w.Write([]byte(fmt.Sprintf("Authentication unsuccessful, please check password or email: %s", email)))
}
