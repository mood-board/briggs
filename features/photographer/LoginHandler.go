package photographer

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/ofonimefrancis/brigg/internal/config"
	"github.com/ofonimefrancis/brigg/message"

	"gopkg.in/mgo.v2/bson"
)

type JwtToken struct {
	Token string `json:"token"`
}

//LoginHandler Login a User and generates a JWT token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userCredential Credentials
	var user Photographer
	if err := json.NewDecoder(r.Body).Decode(&userCredential); err != nil {
		log.Println(message.InvalidPayload)
		return
	}

	//Validate User details
	session := config.Get().Session.Copy()
	defer r.Body.Close()
	defer session.Close()

	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)

	query := collection.Find(bson.M{"email": userCredential.Email})
	if err := query.One(&user); err != nil {
		log.Println("Account does not exist")
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(userCredential.Password)); err != nil {
		log.Println("Account exist... Creating token")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":    userCredential.Email,
			"password": userCredential.Password,
		})
		tokenString, _ := token.SignedString([]byte(config.Get().JwtKey))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	}

}
