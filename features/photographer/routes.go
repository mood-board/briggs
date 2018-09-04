package photographer

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/ofonimefrancis/brigg/internal/config"
	"github.com/ofonimefrancis/brigg/message"
)

//Routes Photographer route
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Put("/upload/{user_id}", UploadAvatar)
	//CRUD
	router.Get("/{username}", FindUser)
	router.Get("/", Find)
	router.Get("/id/{user_id}", GetUserByID)
	router.Post("/authenticate/signup", SignUpHandler)
	router.Post("/authenticate/login", LoginHandler)
	return router
}

//UploadAvatar Upload(Changes the Users Avatar URL and uploads to CloudStorage)
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	var p Photographer
	user, err := p.FindByID(config.Get(), userID)
	if err != nil {
		log.Printf("Invalid UserID: %v\n", err)
		message.NewAPIError(&message.APIError{Message: "Invalid User ID"}, w)
		return
	}

	imageURL, err := uploadProfilePicture(r)
	if err != nil {
		log.Printf("Error Uploading Profile Picture%v\n", err)
		message.NewAPIError(&message.APIError{Message: "Error Uploading Profile Picture"}, w)
		return
	}

	fmt.Printf("Image URL: %s\n", imageURL)

	user.AvatarURL = imageURL
	user.UpdatedAt = time.Now()

	if err := user.Update(config.Get(), user); err != nil {
		log.Printf("Error Connecting to database... %v\n", err)
		message.NewAPIError(&message.APIError{Message: "Database Error"}, w)
		return
	}

	log.Println("All done... Refactor codebase")
	message.NewAPIResponse(&message.APIResponse{Success: true, Message: "Upload Successful"}, w, http.StatusOK)
	return
}
