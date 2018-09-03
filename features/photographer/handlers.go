package photographer

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

//GetUserByID Retrieves a user using their ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	fmt.Println("param: ", userID)
	var user Photographer
	p, err := user.FindByID(config.Get(), userID)
	if err != nil {
		log.Println("Error retrieving user data", err)
		message.NewAPIError(&message.APIError{Message: "Error retrieving user data"}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Data: p, Success: true}, w, http.StatusOK)
	return
}

//UpdateAvatar PUT request to update the users avatar URL
func UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	imageURL, err := uploadProfilePicture(r)
	if err != nil {
		//message.NewAPIError(&message.APIError{Message: "Error Uploading avatar"}, w)
		log.Printf("Error uploading image %v", err)
		return
	}
	var p Photographer
	user, err := p.FindByID(config.Get(), userID)
	if err != nil {
		message.NewAPIError(&message.APIError{Message: "Invalid User ID"}, w)
		log.Printf("Invalid User ID: %v", err)
		return
	}
	user.AvatarURL = imageURL
	user.UpdatedAt = time.Now()

	if err := p.Update(config.Get(), user); err != nil {
		message.NewAPIError(&message.APIError{Message: "Error Updating Database"}, w)
		log.Printf("Error Updating Database: %v", err)
		return
	}

	message.NewAPIResponse(&message.APIResponse{Success: true, Message: "Profile Picture updated"}, w, http.StatusOK)
	return
}

func Find(w http.ResponseWriter, r *http.Request) {
	var user Photographer
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 1
	}

	p, err := user.Listings(config.Get(), page)
	if err != nil {
		log.Println(err)
		message.NewAPIError(&message.APIError{Message: "Error retrieving Listing"}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Data: p}, w, http.StatusOK)
	return
}
