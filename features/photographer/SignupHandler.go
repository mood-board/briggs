package photographer

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"

	"github.com/ofonimefrancis/brigg/internal/config"
	"github.com/ofonimefrancis/brigg/message"
	"github.com/ofonimefrancis/brigg/utils"
)

type UserAuthResponse struct {
	User    Photographer `json:"user"`
	Token   string       `json:"token"`
	Message string       `json:"message"`
}

// SignUpHandler is used to handler signup via various approaches
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var photographer Photographer
	err := json.NewDecoder(r.Body).Decode(&photographer)
	if err != nil {
		message.NewAPIError(&message.APIError{Message: message.InvalidPayload}, w)
		log.Println(message.InvalidPayload)
		return
	}
	defer r.Body.Close()

	if photographer.Exists(config.Get()) {
		message.NewAPIError(&message.APIError{Message: message.UserAlreadyExists}, w)
		log.Println(message.UserAlreadyExists)
		return
	}

	uid, _ := uuid.NewV4()

	pwdHash, _ := utils.HashPassword(photographer.Password)
	photographer.HashedPassword = pwdHash
	photographer.Password = ""

	photographer.ID = uid.String()
	photographer.IsActive = true
	photographer.CreatedAt = time.Now()
	photographer.UpdatedAt = time.Now()

	photographer.Add(config.Get())

	//TODO: Send E-mail

	//Generate JWT token
	_, tokenString, err := tokenAuth.Encode(jwtauth.Claims{"id": photographer.ID, "email": photographer.Email})
	if err != nil {
		message.NewAPIError(&message.APIError{Status: http.StatusInternalServerError, Message: "Internal Server Error"}, w)
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
		User:    photographer,
		Message: "Signup success",
		Token:   tokenString,
	}
	render.JSON(w, r, resp)
}
