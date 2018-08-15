package photographer

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
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
		log.Println("password err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, "Internal Server Error")
		return

	}

	_, tokenString, err := tokenAuth.Encode(jwtauth.Claims{"email": user.Email, "id": user.ID})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, "Internal Server Error")
		return
	}

	expires := time.Now().AddDate(1, 0, 0)

	ck := http.Cookie{
		Name:     "jwt",
		HttpOnly: false,
		Path:     "/",
		Expires:  expires,
		Value:    tokenString,
	}

	// write the cookie to response
	http.SetCookie(w, &ck)

	resp := UserAuthResponse{
		User:    user,
		Message: "Login success",
		Token:   tokenString,
	}

	log.Printf("login %#v", resp)
	render.JSON(w, r, resp)
}
