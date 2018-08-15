package tags

import "github.com/ofonimefrancis/brigg/internal/config"

//Tag represents the tag on an image
type Tag struct {
	Text string `json:"tag"`
}

//Add Adds a new Tag
func Add(tagName string) error {
	session := config.Get().Session.Copy()
	defer session.Close()
	var tag Tag
	tag.Text = tagName
	collection := session.DB(config.DATABASENAME).C(config.TAGCOLLECTION)
	return collection.Insert(tag)
}
