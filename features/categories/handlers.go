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
