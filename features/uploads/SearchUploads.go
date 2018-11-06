package uploads

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/brigg/internal/config"
)

func Search(w http.ResponseWriter, r *http.Request) {

}

//FindByTag Searches for uploads that have a certain tag name
func FindByTag(tag, uploadType string, page int) error {
	session := config.Get().Session.Copy()
	defer session.Close()
	pageSize := 20
	skips := pageSize * (page - 1)
	var allUploads []Uploads
	collection := session.DB(config.DATABASENAME).C(config.UPLOADSCOLLECTION)
	return collection.Find(bson.M{"type": uploadType, "tags": bson.M{"$in": tag}}).Sort("-created_at").Skip(skips).Limit(pageSize).All(&allUploads)
}
