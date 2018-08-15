package uploads

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ofonimefrancis/brigg/message"
	"github.com/ofonimefrancis/brigg/utils"
	uuid "github.com/satori/go.uuid"
	storage "google.golang.org/api/storage/v1"
)

type UploadResponse struct {
	UserID string `json:"user_id"`
}

//UploadHandler Uploads the file via http POST request
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: Loo
	userID := "f9bd707b-7a86-4605-8daa-fb50d85dfb9b"
	err := FileUpload(r, userID)
	if err != nil {
		log.Println("Error uploading file ", err)
		message.NewAPIError(&message.APIError{Message: "Error Uploading File to server..."}, w)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Success: true, Message: "Success Uploading file"}, w, http.StatusMultiStatus)

}

//FileUpload Uploads any file for now
//TODO: Specifically upload images and videos, all formats will be discarded
//MaxFile size for images 4MB
func FileUpload(r *http.Request, userID string) error {
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("upload")
	defer file.Close()

	if err != nil {
		return err
	}

	return uploadToGoogleCloud(file, handler.Filename)
}

func uploadToGoogleCloud(file io.Reader, fileName string) error {
	service, err := utils.GetGoogleCloud()
	if err != nil {
		log.Println("Error Connecting to google cloud")
		return errors.New("Error Connecting to google cloud")
	}

	uid, _ := uuid.NewV4()
	fileName = fmt.Sprintf("seemars_%s-%s", uid.String(), strings.ToLower(fileName))

	object := &storage.Object{
		Name:         fileName,
		CacheControl: "public, max-age=31536000",
	}

	_, err = service.Objects.Insert("yescort", object).Media(file).Do()
	if err != nil {
		return err
	}

	return nil
}
