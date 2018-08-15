package uploads

import (
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/new", UploadHandler)
	return router
}
