package photographer

import (
	"fmt"

	"github.com/go-chi/jwtauth"
	"github.com/ofonimefrancis/brigg/internal/config"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", config.Get().Encryption.Private, config.Get().Encryption.Public)
	_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func GetTokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}
