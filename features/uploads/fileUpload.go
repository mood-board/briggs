package uploads

import (
	"fmt"
	"log"
	"net/http"
	"path"

	uuid "github.com/satori/go.uuid"
	storage "google.golang.org/api/storage/v1"
)

func uploadFileFromForm(r *http.Request) (url string, err error) {
	file, fh, err := r.FormFile("uploads")
	if err == http.ErrMissingFile {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	// random filename, retaining existing extension.
	name := fmt.Sprintf("%s_%s%s", uuid.Must(uuid.NewV4()).String(), fh.Filename, path.Ext(fh.Filename))

	service, err := getGoogleCloud()
	if err != nil {
		log.Println(err)
	}

	object := &storage.Object{
		Name:         name,
		CacheControl: "public, max-age=31536000",
		Acl:          []*storage.ObjectAccessControl{{Entity: "allUsers", Role: "READER"}},
	}
	w, err := service.Objects.Insert("yescort", object).Media(file).Do()
	if err != nil {
		return "", err
	}
	fmt.Printf("%v", w)
	//w.Acl = []*storage.ObjectAccessControl{{Entity: "allUsers", Role: "READER"}}

	const publicURL = "https://storage.googleapis.com/%s/%s"
	return fmt.Sprintf(publicURL, "yescort", name), nil
}
