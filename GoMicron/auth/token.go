// token.go
package auth

import (
	"encoding/base64"
	"math/rand"
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