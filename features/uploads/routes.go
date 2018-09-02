package uploads

import (
	"github.com/go-chi/chi"
)

var (
	PHOTO = "photo"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", PhotoUploadHandler)
	router.Post("/new/{userID}", UploadHandler)
	return router
}
