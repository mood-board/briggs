package categories

import "time"

//TODO: Work on a flat category, where there is no sub-categories
//Admin will have to create categories and add user uploads to a certain category
type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"photographer_id"`
	ImageURL    string    `json:"imageUrl"`
	DateAdded   time.Time `json:"dateAdded"`
}
