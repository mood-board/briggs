package photographer

import (
	"github.com/go-chi/chi"
)

//Routes Photographer route
func Routes() *chi.Mux {
	router := chi.NewRouter()

	//CRUD
	router.Get("/{username}", FindUser)
	router.Get("/", Find)
	router.Post("/authenticate/signup", SignUpHandler)
	router.Post("/authenticate/login", LoginHandler)
	return router
}