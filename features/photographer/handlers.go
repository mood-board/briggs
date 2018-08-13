package photographer

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ofonimefrancis/brigg/internal/config"
	"github.com/ofonimefrancis/brigg/message"
)

const (
	APPSECRET = "7ae209d0af5791ca5d463a51fa95e62e"
	APPID     = "292483241510833"
)

func FindUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	fmt.Println(username)
	var user Photographer
	p, err := user.Find(config.Get(), username)
	if err != nil {
		log.Println("Error retrieving user data", err)
		message.NewAPIError(&message.APIError{Message: "Error retrieving user data"}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Data: p}, w, http.StatusOK)
	return
}
