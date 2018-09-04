package photographer

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/ofonimefrancis/brigg/internal/config"
)

//Routes Photographer route
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/upload/{user_id}", UploadAvatar)
	//CRUD
	router.Get("/{username}", FindUser)
	router.Get("/", Find)
	router.Get("/id/{user_id}", GetUserByID)
	router.Put("/avatar/{user_id}", UpdateAvatar)
	router.Post("/authenticate/signup", SignUpHandler)
	router.Post("/authenticate/login", LoginHandler)
	return router
}

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	var p Photographer
	user, err := p.FindByID(config.Get(), userID)
	if err != nil {
		log.Printf("Invalid UserID: %v\n", err)
		return
	}

	imageURL, err := uploadProfilePicture(r)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	fmt.Printf("Image URL: %s\n", imageURL)

	// user.Update(config.Get(), bson.M{"$set": bson.M{
	// 	"id":            user.ID,
	// 	"username":      user.Username,
	// 	"first_name":    user.FirstName,
	// 	"password":      user.Password,
	// 	"hash_password": user.HashedPassword,
	// 	"email":         user.Email,
	// 	"last_name":     user.LastName,
	// 	"avatar_url":    imageURL,
	// 	"city":          user.City,
	// 	"country":       user.Country,
	// 	"updated_at":    time.Now(),
	// }})

	user.AvatarURL = imageURL
	user.UpdatedAt = time.Now()

	user.Update(config.Get(), user)
	log.Println("All done... Refactor codebase")
}
