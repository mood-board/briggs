package categories

import (
	"github.com/go-chi/chi"
)

//Routes Category Routes
func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", AllCategories)
	r.Post("/", AddCategory)
	return r
}
