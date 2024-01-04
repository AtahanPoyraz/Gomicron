// token.go
package auth

import (
	"encoding/base64"
	"math/rand"
	"time"

	protoc"github.com/AtahanPoyraz/protoc"

	"github.com/dgrijalva/jwt-go"
)

//---[ Generate Token ]-------------------------------------------------------------------------------------------------------------------------------//

func init() {
	globalToken = generateToken()
}

func GenerateToken() *Token {
	return &Token{
		AuthToken: globalToken.AuthToken,
	}
}

func generateToken() *Token {
	tokenLength := 32 

	randomBytes := make([]byte, tokenLength)
	rand.Read(randomBytes)

	token := base64.URLEncoding.EncodeToString(randomBytes)

	return &Token{
		AuthToken: token,
	}
}

func GenerateJWTToken(id int, email, password string) (*Token, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	var jwtKey = []byte("my_secret_key")
	claims := &Claims{
		Users: protoc.Users{
            Id:       int64(id),
            Email:    email,
            Password: password,
        },
		StandardClaims: jwt.StandardClaims{
	  		ExpiresAt: expirationTime.Unix(),
	 },
	}
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
	 return nil, err
	}

    return &Token{
		AuthToken: tokenString,
	}, nil
}