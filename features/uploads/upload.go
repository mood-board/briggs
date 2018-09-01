package uploads

import (
	"fmt"
	"io"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
)

func MultiUploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(200000)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formData := r.MultipartForm
	files := formData.File["uploads"]

	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		uid, _ := uuid.NewV4()
		outputFile, err := os.Create(fmt.Sprintf("/tmp/%s_%s", uid.String(), files[i].Filename))
		defer outputFile.Close()

		if err != nil {
			fmt.Fprintln(w, "Unable to create the file for writing. Check your write access priviledges.")
			return
		}
		_, err = io.Copy(outputFile, file)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Fprintf(w, "Files uploaded successfully")
	}
}
