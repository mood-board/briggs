package uploads

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/brigg/internal/config"
)

//Uploads Images uploaded by photographers
type Uploads struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Title      string    `json:"title"`
	URL        string    `json:"upload_url"` //URL to file on google cloudstorage
	UploadedAt time.Time `json:"created_at"`
	Tags       string    `json:"tags"`
	IsFeatured bool      `json:"featured"`
	Likes      int       `json:"likes"`
	Favorites  int       `json:"favorites"`
	Comments   int       `json:"comments"`
	Download   int       `json:"downloads"`
	UsersName  string    `json:"user"`
	Type       string    `json:"type"`
	PageURL    string    `json:"page_url"`
	ImageWidth string    `json:"width"`
	ImageSize  int       `json:"image_size"`
}

//Add Saves an upload with meta data to the database
func (u *Uploads) Add(conf *config.Config) error {
	session := conf.Session.Copy()
	defer session.Close()
	collection := session.DB(config.DATABASENAME).C(config.UPLOADSCOLLECTION)
	return collection.Insert(u)
}

//Listings Fetch Image Uploads from a USer
func (u *Uploads) Listings(conf *config.Config, userID string, page int) ([]Uploads, error) {
	session := conf.Session.Copy()
	defer session.Close()

	var uploads []Uploads
	pageSize := 20
	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)
	err := collection.Find(bson.M{"user_id": userID}).Skip(pageSize * (page - 1)).Limit(pageSize).All(&uploads)
	return uploads, err
}
