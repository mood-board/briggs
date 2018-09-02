package uploads

import (
	"net/http"
	"strconv"
	"time"

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

	message.NewAPIResponse(&message.APIResponse{Success: true, Data: uploads}, w, http.StatusOK)
	return
}
