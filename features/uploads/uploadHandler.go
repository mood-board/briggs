package uploads

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"github.com/ofonimefrancis/brigg/features/photographer"
	"github.com/ofonimefrancis/brigg/internal/config"
	"github.com/ofonimefrancis/brigg/message"
	uuid "github.com/satori/go.uuid"
)

//PhotoUploadHandler Handles Image Upload by a User
func PhotoUploadHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")
	tags := r.FormValue("tags")
	size, _ := strconv.Atoi(r.FormValue("size"))

	imageURL, err := uploadFileFromForm(r)

	if err != nil {
		message.NewAPIError(&message.APIError{Message: "Error Uploading Image to storage"}, w)
		return
	}

	var u photographer.Photographer

	user, err := u.FindByID(config.Get(), userID)
	if err != nil {
		message.NewAPIError(&message.APIError{Message: "Cannot upload Image not tied to user"}, w)
		return
	}

	uploads := Uploads{
		ID:         uuid.Must(uuid.NewV4()).String(),
		UserID:     userID,
		URL:        imageURL,
		UploadedAt: time.Now(),
		UsersName:  user.FirstName,
		Type:       PHOTO, //TODO: Check the uploaded type and assign type based on that
		Tags:       tags,
		ImageSize:  size,
	}
	if err := uploads.Add(config.Get()); err != nil {
		log.Printf("Error saving upload to database: %v", err)
		message.NewAPIError(&message.APIError{Message: "Error uploading to databae"}, w)
	}

	message.NewAPIResponse(&message.APIResponse{Success: true, Data: uploads}, w, http.StatusOK)
	return
}

//MyImages Retrieves User Uploaded Images
func MyImages(w http.ResponseWriter, r *http.Request) {
	//Get all the images by the user
	userID := chi.URLParam(r, "user_id")
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 1
	}

	var uploads Uploads

	allImages, err := uploads.Listings(config.Get(), userID, page)

	if err != nil {
		log.Printf("Error retrieving uploads from the database: %v", err)
		message.NewAPIError(&message.APIError{Message: "Error retrieving uploads from the database"}, w)
		return
	}

	message.NewAPIResponse(&message.APIResponse{Success: true, Data: allImages}, w, http.StatusOK)
	return
}
