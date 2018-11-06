package categories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/brigg/internal/config"
)

func (c *Category) Add(conf *config.Config) error {
	session := conf.Session.Copy()
	defer session.Close()
	collection := session.DB(config.DATABASENAME).C(config.CATEGORIESCOLLECTION)
	return collection.Insert(c)
}

//TODO: Add Pagination and sort by date created
func (c *Category) FindByCategoryName(conf *config.Config, categoryName string) ([]Category, error) {
	session := conf.Session.Copy()
	defer session.Close()
	var uploadsInCategory []Category
	collection := session.DB(config.DATABASENAME).C(config.CATEGORIESCOLLECTION)
	err := collection.Find(bson.M{"name": categoryName}).All(&uploadsInCategory)
	if err != nil {
		return nil, err
	}
	return uploadsInCategory, nil
}

//ListCategories A list of all the categories
func (c *Category) ListCategories(conf *config.Config) ([]Category, error) {
	session := conf.Session.Copy()
	defer session.Close()
	var allCategories []Category
	collection := session.DB(config.DATABASENAME).C(config.CATEGORIESCOLLECTION)
	if err := collection.Find(bson.M{}).All(&allCategories); err != nil {
		return nil, err
	}
	return allCategories, nil
}
