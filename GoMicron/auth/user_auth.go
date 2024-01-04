package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AtahanPoyraz/cmd"
	"github.com/AtahanPoyraz/config"
	protoc "github.com/AtahanPoyraz/protoc"
	"github.com/AtahanPoyraz/db"
)

//---[ USERS DB CONNECT ]----------------------------------------------------------------------------------------------------------------------------//

func init() {
	c, err := config.ReadConfigFromFile("./config.yml")
	if err != nil {
		l.Printf("%s[ERROR]%s : Read operation failed please check file : %v", cmd.BRED_WHITE, cmd.TRESET, err)
		os.Exit(1)
	}
	conf = c
	pqdb = db.NewDB(conf.Services.Postgres.DBHost, conf.Services.Postgres.DBPort, conf.Services.Postgres.DBUser,
		conf.Services.Postgres.DBPass, conf.Services.Postgres.DBName, conf.Services.Postgres.DBSSLMode,
		conf.Services.Postgres.DBNET, conf.Services.Postgres.DBTimeout)
}

//---[ USER AUTH ]-----------------------------------------------------------------------------------------------------------------------------------//

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User_List []*protoc.Users

func (s *Server) AuthUser(w http.ResponseWriter, r *http.Request) {
	async.Add(1)
	go pqdb.OpenDB(conf.Services.Postgres.DBSrvname, channel)
	result := <-channel

	if result.Err != nil {
		l.Printf("%v", result.Err)
	}

	db := result.DB

	var users User_List

	rows, err := db.Query("SELECT id, email, password FROM users")
	if err != nil {
		l.Printf("failed to execute query: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user protoc.Users
		err = rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			l.Printf("failed to scan row: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		users = append(users, &user)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials Credentials
	err = json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Error decoding JSON request body", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if credentials.Email == user.Email && credentials.Password == user.Password {
			token := GenerateToken()
			response := map[string]interface{}{
				"message":       "Authentication successful",
				"authenticated": true,
				"token":         token.AuthToken,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

			s.l.Printf("%s[AUTH]%s : Authentication successful for user: %s", cmd.BBLACK_VIOLET, cmd.TRESET, credentials.Email)
			return
		}
	}

	response := map[string]interface{}{
		"message":       "Authentication unsuccessful",
		"authenticated": false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	w.Write([]byte(fmt.Sprintf("Authentication unsuccessful, please check password or email: %s", credentials.Email)))
}