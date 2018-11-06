package categories

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ofonimefrancis/brigg/internal/config"
	"github.com/ofonimefrancis/brigg/message"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		log.Printf("Error decoding payload: %v\n", err)
		message.NewAPIError(&message.APIError{Message: "Invalid Payload"}, w)
		return
	}

	if err := newCategory.Add(config.Get()); err != nil {
		log.Printf("Error creating a new Category: %v\n", err)
		message.NewAPIError(&message.APIError{Message: "Error creating a new category"}, w)
		return
	}

	message.NewAPIResponse(&message.APIResponse{Success: true, Message: "New Category successfully created"}, w, http.StatusCreated)
}

func AllCategories(w http.ResponseWriter, r *http.Request) {
	category := new(Category)
	allCategories, err := category.ListCategories(config.Get())
	if err != nil {
		log.Printf("AllCategories-Error %v\n", err)
		message.NewAPIError(&message.APIError{Message: "Error retrieving categories"}, w)
		return
	}
	if len(allCategories) == 0 {
		log.Printf("There are currently no categories in our database")
		message.NewAPIResponse(&message.APIResponse{Data: map[string]string{}}, w, http.StatusOK)
		return
	}
	message.NewAPIResponse(&message.APIResponse{Data: allCategories}, w, http.StatusOK)
}
