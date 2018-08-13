package photographer

import (
	"github.com/go-chi/chi"
)

func Route() *chi.Mux {
	router := chi.NewRouter()

	//Get the Users Profile
	//?user/$username
	router.Get("/{username}", nil)

	return router
}
