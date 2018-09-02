package uploads

import (
	"net/http"
	"time"

	"github.com/ofonimefrancis/brigg/features/photographer"
	"github.com/ofonimefrancis/brigg/internal/config"
	"github.com/ofonimefrancis/brigg/message"
	uuid "github.com/satori/go.uuid"
)

func PhotoUploadHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")
	tags := r.FormValue("tags")
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
		Type:       PHOTO,
		Tags:       tags,
	}

	message.NewAPIResponse(&message.APIResponse{Success: true, Data: uploads}, w, http.StatusOK)
	return

}
