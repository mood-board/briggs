package photographer

import (
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	//Get the Users Profile
	//?user/$username

	//CRUD
	router.Get("/{username}", FindUser)
	router.Post("/", SignUpHandler)
	router.Post("/authenticate", LoginHandler)
	return router
}
