package uploads

import (
	"github.com/ofonimefrancis/brigg/features/tags"
	"github.com/ofonimefrancis/brigg/internal/config"
)

//Uploads Images uploaded by photographers
type Uploads struct {
	ID         string     `json:"id"`
	UserID     string     `json:"user_id"`
	Title      string     `json:"title"`
	URL        string     `json:"upload_url"` //URL to file on google cloudstorage
	UploadedAt string     `json:"created_at"`
	Tags       []tags.Tag `json:"tags"`
	IsFeatured bool       `json:"featured"`
}

//Add Saves an upload with meta data to the database
func (u *Uploads) Add(conf *config.Config) error {
	session := conf.Session.Copy()
	defer session.Close()
	collection := session.DB(config.DATABASENAME).C(config.UPLOADSCOLLECTION)
	return collection.Insert(u)
}
